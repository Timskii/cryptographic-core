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

	transactions := []*protos.Transaction{}

	ledger1 := ledger.InitTestLedger()
	if ledger1.GetBlockchainSize() == 0 {
		if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
			makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
			fmt.Printf("makeGenesisError= %+v\n", makeGenesisError)
		}
	}
	
	for i := 0; i < len(jsonobject); i++{
		ledger1.BeginTxBatch(1)
		args := jsonobject[i].Params.CtorMsg.Args
		fmt.Println("\n\n------------transaction --------------")
		fmt.Printf("jsonobject: %+v\n", jsonobject[i])
		transaction, err := protos.NewTransaction(	protos.ChaincodeID{Name: jsonobject[i].Params.ChaincodeID.Name}, 
													ut.GenerateUUID(), 
													jsonobject[i].Params.CtorMsg.Function, 
													args)
		if err != nil {
			fmt.Printf("Error creating NewTransaction: %s", err)
		}

		//if i == 0 {
			ledger1.TxBegin(transaction.Txid)
		//} 

		ledger1.SetState(jsonobject[i].Params.ChaincodeID.Name, util.GenerateKey(&args), []byte(args[2]+args[3]))
		
		transactions = append(transactions,transaction)
		//transactions[0] = transaction
		//if i == len(jsonobject)-1 {
			ledger1.TxFinished(transactions[i].Txid, true)
		//}
		ledger1.CommitTxBatch(1, transactions, nil, []byte("dummy-proof"))	//COMN
	}	

	
	

	fmt.Println((fmt.Sprintf("%+v\n", transactions)))
}