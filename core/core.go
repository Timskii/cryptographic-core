package core

import (
		"fmt"
		"strconv"
		"encoding/json"
		"encoding/base64"
		"bytes"
		"io/ioutil"
		"os"

		"github.com/hyperledger/fabric/util"
		"github.com/hyperledger/fabric/core/ledger"
		ut "github.com/hyperledger/fabric/core/util"
		"github.com/hyperledger/fabric/protos"
		"github.com/hyperledger/fabric/core/ledger/statemgmt"
		"github.com/hyperledger/fabric/core/db"
)

//Функция для добавления данных из файла в базу данных
func AddData (jsonobject []*Jsonobject){
	var commited bool = false
	var transactions []*protos.Transaction
	var oldValueByte1,oldValueByte2 []byte
	var finalValue1,finalValue2 int

	ledger1,_ := ledger.GetLedger()
	blockchainSize := ledger1.GetBlockchainSize()
	if blockchainSize == 0 {
		panic("!Внимание!\nНе прошла инициализация блока, запустите приложение с аргументом 'i'")
	}
	ledger1.BeginTxBatch(blockchainSize)

	for i := 0; i < len(jsonobject); i++{
		if i==0 {commited = true}else{commited = false}
		args := jsonobject[i].Params.CtorMsg.Args
		chaincodeID := jsonobject[i].Params.ChaincodeID.Name
		transaction,err := createTransaction(args,chaincodeID)
		if err != nil {
			fmt.Println("ошибка создания транзакции: ",err)
		}else {
			transactions = append(transactions,transaction)
			ledger1.TxBegin(transactions[i].Txid)
			if jsonobject[i].Params.CtorMsg.Function == "register" {
				ledger1.SetState(chaincodeID, util.GenerateKey(args[0], args[1]), []byte(args[2]))
			} else {
				key1 := util.GenerateKey(args[0], args[1])
				key2 := util.GenerateKey(args[0], args[2])
				oldValueByte1, _ = ledger1.GetState(chaincodeID, key1, commited)
				oldValueByte2, _ = ledger1.GetState(chaincodeID, key2, commited)
				sum, _ := strconv.Atoi(args[3])
				oldValue1, _ := strconv.Atoi(string(oldValueByte1))
				oldValue2, _ := strconv.Atoi(string(oldValueByte2))
				finalValue1 = oldValue1 - sum
				finalValue2 = oldValue2 + sum
				ledger1.SetState(chaincodeID, key1, []byte(strconv.Itoa(finalValue1)))
				ledger1.SetState(chaincodeID, key2, []byte(strconv.Itoa(finalValue2)))
			}
			ledger1.TxFinished(transactions[i].Txid, true)
		}
	}
	if len(transactions)>0 {
		ledger1.CommitTxBatch(blockchainSize, transactions, nil, nil)
		fmt.Printf("Запись блока прошла успешно, %d транзакций обработано.", len(transactions))
	}
}

// Функция для формирования транзакции
func createTransaction(args []string, chaincodeID string) (*protos.Transaction, error) {
	transaction, err := protos.NewChaincodeExecute(
		&protos.ChaincodeInvocationSpec{
			ChaincodeSpec: &protos.ChaincodeSpec{
				CtorMsg: &protos.ChaincodeInput{
					Args: util.ToChaincodeArgs(args)},
				ChaincodeID: &protos.ChaincodeID{Name: chaincodeID},
			}},
		ut.GenerateUUID(),
		protos.Transaction_CHAINCODE_INVOKE)
	if err != nil {
		return nil,err
	}
	return transaction,nil
}

func CreateNilBlock(){
	ledger1,_ := ledger.GetLedger()
	if ledger1.GetBlockchainSize() != 0 {
		fmt.Printf("Внимание!\nИнициализация блока уже прошла!")
	}else{
		if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
			makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
			if makeGenesisError != nil {
				fmt.Printf("Внимание! во время инициализации блока возникла ошибка: %+v\n", makeGenesisError)
			}else{
				fmt.Println("Инициализация базы данных прошла успешно!")
			}
		}
	}
}

const level  = 9

// фнкция по проверке транзакции
func ReadTransaction(idx string){
	var payload Payload
	var chaincodeID		ChaincodeID
	var hash,hashDB,state []byte
	var bucketKey *bucketKey
	var isValid bool = true

	ledger,_ := ledger.GetLedger()
	transaction, _ := ledger.GetTransactionByID(idx)
	if transaction != nil {
		fmt.Printf("Транзакция с id <%s> = %s\n",idx,transaction)
		blockNumber := util.GetBlockNumberByTransaction(idx)

		json.Unmarshal(transaction.Payload, &payload)
		json.Unmarshal(transaction.ChaincodeID, &chaincodeID)

		args := payload.ChaincodeSpec.CtorMsg.Args
		for i, arg := range args {
			str, _ := base64.StdEncoding.DecodeString(arg)
			args[i] = string(str)
		}

		for i:=1; i<len(args)-1; i++ {
			state, bucketKey = getState(args[0], args[i], chaincodeID, blockNumber)
			hash=ut.ComputeCryptoHash(state)
			hashDB = getHashFromDB(bucketKey, blockNumber)
			fmt.Printf(	"\nСостояние %d   = %x\n"+
								"Хеш состояния = %x\n"+
								"Хеш из базы   = %x\n", i,state, hash, hashDB)
			if bytes.Compare(hashDB , hash) !=0 {isValid = false}
		}
		if isValid == true {
			fmt.Println("Транзакция корректна.")
		}else{
			fmt.Println("Транзакция не корректна!")
		}
	}else{
		fmt.Printf("Транзакция с id <%s> не найдена!",idx)
	}

}

//Функция по получению состояния
func getState(nin string, user string,chaincodeID ChaincodeID, blockNumber int) ([]byte, *bucketKey){
	key := util.GenerateKey(nin,user)
	hashFunction := fnvHash
	conf := &config{hashFunc:hashFunction}
	compositeKey := statemgmt.ConstructCompositeKey(chaincodeID.Name, key)
	bucketHash := conf.computeBucketHash(compositeKey)
	bucketNumber := int(bucketHash)%1000003 + 1
	bucketKey := &bucketKey{level, bucketNumber}
	dataKey	:= &dataKey{bucketKey,compositeKey}
	openchainDB := db.GetDBHandle()
	nodeBytes, _ := openchainDB.GetFromStateCFForBlockNumber(dataKey.getEncodedBytes(),blockNumber)

	state := &hash{nil}
	state.appendSizeAndData([]byte(chaincodeID.Name))
	state.appendSize(1)
	state.appendSizeAndData([]byte(key))
	state.appendSizeAndData(nodeBytes)

	return state.hashingData,bucketKey
}

// Функция по получению хэша состояния из базы данных
func getHashFromDB (bucketKey *bucketKey, blockNumber int) []byte{
	openchainDB := db.GetDBHandle()
	bucketKey.level = level-1
	bucketKey.bucketNumber = computeBucketNumber(bucketKey.bucketNumber)
	hashDBKey := bucketKey.getEncodedBytes()
	hashDB ,_ := openchainDB.GetFromStateCFForBlockNumber(hashDBKey,blockNumber)
	return unmarshalCryptoHash(hashDB)
}

//Функция по проверке блоков в базе данных
func TestValidAllBlocks(){
	var previousBlockHash []byte
	ledger,_ := ledger.GetLedger()
	var countErrBlocks int =0
	size := ledger.GetBlockchainSize()

	blockNil := new(protos.Block)
	blockNil.NonHashData = &protos.NonHashData{LocalLedgerCommitTimestamp: ut.CreateUtcTimestamp()}

	for i:= 1; i<int(size); i++{
		block, _ := ledger.GetBlockByNumber(uint64(i))
		if i == 1{
			previousBlockHash,_ = blockNil.GetHash()
		}else {
			previousBlock, _ := ledger.GetBlockByNumber(uint64(i - 1))
			previousBlockHash, _ = previousBlock.GetHash()
		}
		if block == nil || !bytes.Equal(previousBlockHash,block.PreviousBlockHash){
			countErrBlocks++
			fmt.Printf("E")
		} else {
			fmt.Printf(".")
		}
	}
	if countErrBlocks > 0 {
		fmt.Printf("\nВнимание, выявлено %d невалидных блоков!",countErrBlocks)
	}else{
		fmt.Printf("\nВсе %d блок(-ов) успешно прошли проверку на валидность.", size-1)
	}
}
// Функция по проверке несанкционированного изменения программы
func Checksum(){
	fileData,_ := ioutil.ReadFile("main.exe")
	hashFile := fileData[(len(fileData)-64):]
	fileData = fileData[:(len(fileData)-64)]
	hash := ut.ComputeCryptoHash(fileData)
	if !bytes.Equal(hash,hashFile) {
		fmt.Printf("Внимание, выявлено несанкционированное изменение программы!")
		os.Exit(1)
	}
}