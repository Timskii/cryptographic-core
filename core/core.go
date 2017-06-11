package core

import (
		"fmt"
		//"encoding/json"

		"github.com/hyperledger/fabric/util"
		"github.com/hyperledger/fabric/core/ledger"
		ut "github.com/hyperledger/fabric/core/util"
		"github.com/hyperledger/fabric/protos"
		)


func AddData (jsonobject []*Jsonobject){

	fmt.Printf("AddData \n")
	/*state := state.NewState()
	state.TxFinish("txID", true)*/
	fmt.Printf("Results: %v\n", jsonobject[0])
	fmt.Printf("function: %v\n", jsonobject[0].Params.CtorMsg.Function)

// -----------------------------<Block #0>---------------------
	ledger1 := ledger.InitTestLedger()

	fmt.Printf("ledger1.InitTestLedger: %v\n ", ledger1)
	ledger1.BeginTxBatch(0)
	fmt.Printf("ledger1.BeginTxBatch(0): %v\n ", ledger1)
	err := ledger1.CommitTxBatch(0, []*protos.Transaction{}, nil, []byte("dummy-proof"))
	fmt.Printf("ledger1.CommitTxBatch(0): %v\n ", ledger1)
	if err != nil {
		fmt.Printf("Error in commit: %s", err)
	}

// -----------------------------<Block #1>------------------------------------

	//data, _ := json.Marshal(jsonobject)

	ledger1.BeginTxBatch(1)
	transaction1a, err := protos.NewTransaction(protos.ChaincodeID{Path: "Contracts"}, ut.GenerateUUID(), "NewContract", []string{"name: MyContract1, code: var x; function setX(json) {x = json.x}}"})
	if err != nil {
		fmt.Printf("Error creating NewTransaction: %s", err)
	}

	ledger1.TxBegin(transaction1a.Txid)
	ledger1.SetState("MyContract1", "code", []byte("code example"))
	ledger1.TxFinished(transaction1a.Txid, true)
	ledger1.CommitTxBatch(1, []*protos.Transaction{transaction1a}, nil, []byte("dummy-proof"))


	// -----------------------------<Block #2>------------------------------------

	ledger1.BeginTxBatch(2)
	transaction2a, err := protos.NewTransaction(protos.ChaincodeID{Path: "MyContract"}, ut.GenerateUUID(), "setX", []string{"{x: \"hello\"}"})
	if err != nil {
		fmt.Printf("Error creating NewTransaction: %s", err)
	}
	transaction2b, err := protos.NewTransaction(protos.ChaincodeID{Path: "MyOtherContract"}, ut.GenerateUUID(), "setY", []string{"{y: \"goodbuy\"}"})
	if err != nil {
		fmt.Printf("Error creating NewTransaction: %s", err)
	}

	// Run this transction in the VM. The VM updates the state
	ledger1.TxBegin(transaction2a.Txid)
	ledger1.SetState("MyContract", "x", []byte("hello"))
	ledger1.SetState("MyOtherContract", "y", []byte("goodbuy"))
	ledger1.TxFinished(transaction2a.Txid, true)

	// Commit txbatch that creates the 2nd block on blockchain
	ledger1.CommitTxBatch(2, []*protos.Transaction{transaction2a, transaction2b}, nil, []byte("dummy-proof"))
	// -----------------------------</Block #2>-----------------------------------

	/*transaction := new(Transaction)
	transaction.ChaincodeID = data
	transaction.Txid = "uuid"
	transaction.Timestamp = ut.CreateUtcTimestamp()

	state := buckettree.NewStateImpl()

	configMap := map[string]interface{}{"maxGroupingAtEachLevel":5, "bucketCacheSize":100, "numBuckets":1000003}
	ledger, _ := ledger.GetLedger()
	ledger.GetTempStateHash()

	err2 := state.Initialize(configMap)
	fmt.Printf("Initialize: %v\n ", err2)
	data1, eror := state.Get("ChaincodeID1","keys")
	fmt.Printf("data: %v\n eror: %v\n", data1,eror)*/

	fmt.Printf("transaction2a: %v\n", transaction2a)
	fmt.Printf("transaction2a: %#v\n", transaction2a)

	util.PrintData([]byte(fmt.Sprintf("%v\n", transaction2a)))
}