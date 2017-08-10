package core

import (
		"fmt"
		"strconv"
		//"encoding/json"

		"github.com/hyperledger/fabric/util"
		"github.com/hyperledger/fabric/core/ledger"
		ut "github.com/hyperledger/fabric/core/util"
		"github.com/hyperledger/fabric/protos"
		)


func AddData (jsonobject []*Jsonobject){
	fmt.Printf("AddData \n")
	t := protos.Transaction_CHAINCODE_INVOKE
	transactions := []*protos.Transaction{}

	ledger1,_ := ledger.GetLedger()
	blockchainSize := ledger1.GetBlockchainSize()
	if blockchainSize == 0 {
		panic("!Внимание!\nНепрошла инициализация блока, запустите приложение с аргументом 'i'")
	}

	ledger1.BeginTxBatch(blockchainSize)
	var commited bool = false
	for i := 0; i < len(jsonobject); i++{
		if i==0 {commited = true}else{commited = false}
		args := jsonobject[i].Params.CtorMsg.Args
		chaincodeID := jsonobject[i].Params.ChaincodeID.Name
		if jsonobject[i].Params.CtorMsg.Function == "register" {
			transaction, err := protos.NewChaincodeExecute(
				&protos.ChaincodeInvocationSpec{
					ChaincodeSpec: &protos.ChaincodeSpec{
						CtorMsg: &protos.ChaincodeInput{
							Args: ut.ToChaincodeArgs(args[0], args[1], args[2], )},
						ChaincodeID: &protos.ChaincodeID{Name: chaincodeID},
					}},
				ut.GenerateUUID(),
				t)
			if err != nil {
				fmt.Printf("Error creating NewTransaction: %s", err)
			}
			transactions = append(transactions, transaction)
			ledger1.TxBegin(transactions[i].Txid)
			ledger1.SetState(chaincodeID, util.GenerateKey(args[0],args[1]), []byte(args[2]))
		}else {
			key1 := util.GenerateKey(args[0],args[1])
			key2 := util.GenerateKey(args[0],args[2])
			oldValueByte1, _ := ledger1.GetState(chaincodeID, key1, commited)
			oldValueByte2, _ := ledger1.GetState(chaincodeID, key2, commited)
			fmt.Printf("\ncore.go GetState oldValueByte1 = %s", oldValueByte1)
			fmt.Printf("\ncore.go GetState oldValueByte2 = %s", oldValueByte2)
			value,_ := strconv.Atoi(args[3])
			oldValue1,_ := strconv.Atoi(string(oldValueByte1))
			oldValue2,_ := strconv.Atoi(string(oldValueByte2))

			finalValue1 := oldValue1 - value
			finalValue2 := oldValue2 + value

			transaction, err := protos.NewChaincodeExecute(
				&protos.ChaincodeInvocationSpec{
					ChaincodeSpec: &protos.ChaincodeSpec{
						CtorMsg: &protos.ChaincodeInput{
							Args: ut.ToChaincodeArgs(args[0], args[1], args[2], args[3])},
						ChaincodeID: &protos.ChaincodeID{Name: chaincodeID},
					}},
				ut.GenerateUUID(),
				t)
			if err != nil {
				fmt.Printf("Error creating NewTransaction: %s", err)
			}
			transactions = append(transactions, transaction)
			ledger1.TxBegin(transactions[i].Txid)
			ledger1.SetState(chaincodeID, key1, []byte(strconv.Itoa(finalValue1)))
			ledger1.SetState(chaincodeID, key2, []byte(strconv.Itoa(finalValue2)))
		}
		ledger1.TxFinished(transactions[i].Txid, true)
	}
	ledger1.CommitTxBatch(blockchainSize, transactions, nil, nil)
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