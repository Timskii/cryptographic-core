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


	fmt.Println("\n\n------------InitTestLedger --------------")
	ledger1,_ := ledger.GetLedger() //ledger.InitTestLedger()
	fmt.Println("\n\n------------InitTestLedger --------------")
	if ledger1.GetBlockchainSize() == 0 {
		if makeGenesisError := ledger1.BeginTxBatch(0); makeGenesisError == nil {
			makeGenesisError := ledger1.CommitTxBatch(0, nil, nil, nil)
			fmt.Printf("makeGenesisError= %+v\n", makeGenesisError)
		}
	}
	
	for i := 0; i < len(jsonobject); i++{
		transactions := []*protos.Transaction{}
		ledger1.BeginTxBatch(1)
		args := jsonobject[i].Params.CtorMsg.Args
		fmt.Println("\n\n------------transaction --------------")
		fmt.Printf("\ncore.go jsonobject: %+v\n", jsonobject[i])
		var t protos.Transaction_Type
		t = protos.Transaction_CHAINCODE_INVOKE

		transaction, err := protos.NewChaincodeExecute(
			&protos.ChaincodeInvocationSpec{
					ChaincodeSpec: &protos.ChaincodeSpec{
												CtorMsg: &protos.ChaincodeInput{
														Args: ut.ToChaincodeArgs(args[0],args[1],args[2],args[3])},
												ChaincodeID : &protos.ChaincodeID{Name: jsonobject[i].Params.ChaincodeID.Name},
																}},
			ut.GenerateUUID(),
			t)
		if err != nil {
			fmt.Printf("Error creating NewTransaction: %s", err)
		}
		ledger1.TxBegin(transaction.Txid)
 		ledger1.SetState(jsonobject[i].Params.ChaincodeID.Name, util.GenerateKey(&args), []byte(args[2]+args[3]))
		transactions = append(transactions,transaction)
		ledger1.TxFinished(transaction.Txid, true)
		ledger1.CommitTxBatch(1, transactions, nil, nil)	//COMN
	}	
}