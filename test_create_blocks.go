package main


import(
		"fmt"

		"github.com/hyperledger/fabric/core/ledger"
		"github.com/hyperledger/fabric/core/util"
		"github.com/hyperledger/fabric/protos"
		"github.com/hyperledger/fabric/core/ledger/statemgmt/buckettree"
	
	)


func main (){
	// --------------------------- initit config --------------------------
	state := buckettree.NewStateImpl()
	configMap := map[string]interface{}{"maxGroupingAtEachLevel":5, "bucketCacheSize":100, "numBuckets":1000003}
	err2 := state.Initialize(configMap)
	fmt.Printf("Initialize: %v\n ", err2)


	ledger1 := ledger.InitTestLedger()

	fmt.Printf("ledger1.InitTestLedger: %v\n ", ledger1)
	ledger1.BeginTxBatch(0)
	fmt.Printf("ledger1.BeginTxBatch(0): %v\n ", ledger1)
	err := ledger1.CommitTxBatch(0, []*protos.Transaction{}, nil, []byte("dummy-proof"))
	fmt.Printf("ledger1.CommitTxBatch(0): %v\n ", ledger1)
	if err != nil {
		fmt.Printf("Error in commit: %s", err)
	}


	// -----------------------------<Block #0>---------------------

	// -----------------------------<Block #1>------------------------------------

	// Deploy a contract
	// To deploy a contract, we call the 'NewContract' function in the 'Contracts' contract
	// TODO Use chaincode instead of contract?
	// TODO Two types of transactions. Execute transaction, deploy/delete/update contract
	ledger1.BeginTxBatch(1)
	transaction1a, err := protos.NewTransaction(protos.ChaincodeID{Path: "Contracts"}, generateUUID(), "NewContract", []string{"name: MyContract1, code: var x; function setX(json) {x = json.x}}"})
	if err != nil {
		fmt.Printf("Error creating NewTransaction: %s", err)
	}
	// VM runs transaction1a and updates the global state with the result
	// In this case, the 'Contracts' contract stores 'MyContract1' in its state
	ledger1.TxBegin(transaction1a.Txid)
	ledger1.SetState("MyContract1", "code", []byte("code example"))
	ledger1.TxFinished(transaction1a.Txid, true)
	ledger1.CommitTxBatch(1, []*protos.Transaction{transaction1a}, nil, []byte("dummy-proof"))
	// -----------------------------</Block #1>-----------------------------------

	// -----------------------------<Block #2>------------------------------------

	ledger1.BeginTxBatch(2)
	transaction2a, err := protos.NewTransaction(protos.ChaincodeID{Path: "MyContract"}, generateUUID(), "setX", []string{"{x: \"hello\"}"})
	if err != nil {
		fmt.Printf("Error creating NewTransaction: %s", err)
	}
	transaction2b, err := protos.NewTransaction(protos.ChaincodeID{Path: "MyOtherContract"}, generateUUID(), "setY", []string{"{y: \"goodbuy\"}"})
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
	return
}

func generateUUID() string {
	return util.GenerateUUID()
}
