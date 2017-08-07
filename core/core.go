package core

import (
		"fmt"
		//"strconv"
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
	if ledger1.GetBlockchainSize() == 0 {
		if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
			makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
			fmt.Printf("makeGenesisError= %+v\n", makeGenesisError)
		}
	}

	ledger1.BeginTxBatch(1)
	for i := 0; i < len(jsonobject); i++{
		args := jsonobject[i].Params.CtorMsg.Args
		transaction, err := protos.NewChaincodeExecute(
			&protos.ChaincodeInvocationSpec{
				ChaincodeSpec: &protos.ChaincodeSpec{
					CtorMsg: &protos.ChaincodeInput{
						Args: ut.ToChaincodeArgs(args[0], args[1], args[2], )},
					ChaincodeID: &protos.ChaincodeID{Name: jsonobject[i].Params.ChaincodeID.Name},
				}},
			ut.GenerateUUID(),
			t)
		if err != nil {
			fmt.Printf("Error creating NewTransaction: %s", err)
		}
		transactions = append(transactions, transaction)
		if i == 0 {
			ledger1.TxBegin(transactions[0].Txid)
		}
		ledger1.SetState(jsonobject[i].Params.ChaincodeID.Name, util.GenerateKey(&args), []byte(args[2]))
	}
	ledger1.TxFinished(transactions[0].Txid, true)
	ledger1.CommitTxBatch(1, transactions, nil, nil)
}

func CreateNilBlock(){
	ledger1,_ := ledger.GetLedger()
	if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
		makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
		fmt.Printf("makeGenesisError= %+v\n", makeGenesisError)
	}
}