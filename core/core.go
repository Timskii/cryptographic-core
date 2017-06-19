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
	ledger1.BeginTxBatch(2)
	for i := 0; i < len(jsonobject); i++{
		args := jsonobject[i].Params.CtorMsg.Args
		fmt.Println("\n\n------------transaction --------------")
		fmt.Printf(" %+v\n", jsonobject[i])
		transaction, err := protos.NewTransaction(	protos.ChaincodeID{Name: jsonobject[i].Params.ChaincodeID.Name}, 
													ut.GenerateUUID(), 
													jsonobject[i].Params.CtorMsg.Function, 
													args)
		if err != nil {
			fmt.Printf("Error creating NewTransaction: %s", err)
		}

		if i == 0 {
			ledger1.TxBegin(transaction.Txid)
		} 

		ledger1.SetState(jsonobject[i].Params.ChaincodeID.Name, util.GenerateKey(&args), []byte(args[2]+args[3]))
		
		transactions = append(transactions,transaction)
		if i == len(jsonobject)-1 {
			ledger1.TxFinished(transactions[0].Txid, true)
		} 
	}	

	
	ledger1.CommitTxBatch(2, transactions, nil, []byte("dummy-proof"))	//COMN

	fmt.Println((fmt.Sprintf("%+v\n", transactions)))
}