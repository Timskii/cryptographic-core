package core

import (
		"fmt"
		"strconv"

		"github.com/hyperledger/fabric/util"
		"github.com/hyperledger/fabric/core/ledger"
		ut "github.com/hyperledger/fabric/core/util"
		"github.com/hyperledger/fabric/protos"

	"encoding/json"
	"encoding/base64"
	"github.com/hyperledger/fabric/core/ledger/statemgmt"
	"github.com/hyperledger/fabric/core/db"
)

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
	}
}

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
		panic("!Внимание!\nИнициализация блока уже прошла!")
	}
	if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
		makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
		if makeGenesisError != nil {
			fmt.Printf("Внимание! во время инициализации блока возникла ошибка: %+v\n", makeGenesisError)
		}else{
			fmt.Println("Инициализация блока прошла успешно!")
		}
	}
}

func ReadTransaction(idx string){
	var payload Payload
	var firstKey, secondKey string
	var chaincodeID		ChaincodeID


	ledger,_ := ledger.GetLedger()

	transaction, err := ledger.GetTransactionByID(idx)
	fmt.Printf("\nerr = %s\n", err)

	blockNumber := util.GetBlockNumberByTransaction(idx)

	fmt.Printf("\nblockNumber = %d\n", blockNumber)
	err = json.Unmarshal(transaction.Payload,&payload)
	fmt.Printf("\nerr = %s\n", err)
	err = json.Unmarshal(transaction.ChaincodeID,&chaincodeID)
	fmt.Printf("\nerr = %s\n", err)

	args := payload.ChaincodeSpec.CtorMsg.Args
	for i,arg := range args{
		str,_ := base64.StdEncoding.DecodeString(arg)
		args[i] = string(str)
	}
	fmt.Printf("\nargs = %v\n", args)

	firstKey = util.GenerateKey(args[0],args[1])
	if len(args)> 3{
		secondKey = util.GenerateKey(args[0],args[2])
	}
	fmt.Printf("\nfirstKey = %v\nsecondKey = %v\n", firstKey,secondKey)

	hashFunction := fnvHash
	conf := &config{hashFunc:hashFunction}

	compositeKey1 := statemgmt.ConstructCompositeKey(chaincodeID.Name, firstKey)
	compositeKey2 := statemgmt.ConstructCompositeKey(chaincodeID.Name, secondKey)

	bucketHash := conf.computeBucketHash(compositeKey1)
	bucketHash2 := conf.computeBucketHash(compositeKey2)

	bucketNumber := int(bucketHash)%1000003 + 1
	bucketNumber2 := int(bucketHash2)%1000003 + 1

	dataKey2	:= &dataKey{&bucketKey{9, bucketNumber2},compositeKey2}
	dataKey		:= &dataKey{&bucketKey{9, bucketNumber},compositeKey1}

	openchainDB := db.GetDBHandle()
	nodeBytes, err := openchainDB.GetFromStateCFForBlockNumber(dataKey.getEncodedBytes(),blockNumber)
	nodeBytes2, err := openchainDB.GetFromStateCFForBlockNumber(dataKey2.getEncodedBytes(),blockNumber)

	finalHash := computeCryptoHash(chaincodeID.Name,firstKey,nodeBytes)
	finalHash2 := computeCryptoHash(chaincodeID.Name,secondKey,nodeBytes2)

	fmt.Printf("finalHash = %x\n",finalHash)
	fmt.Printf("finalHash2 = %x\n",finalHash2)


}
