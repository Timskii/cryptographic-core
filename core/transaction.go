package core

//import "fmt"
import "github.com/hyperledger/fabric/core/util"


func NewTransaction(chaincodeID ChaincodeID, uuid string, function string, arguments []string) (*Transaction, error) {
	//data, err := proto.Marshal(&chaincodeID)
	/*if err != nil {
		return nil, fmt.Errorf("Could not marshal chaincode : %s", err)
	}*/
	transaction := new(Transaction)
	//transaction.ChaincodeID = chaincodeID
	transaction.Txid = uuid
	transaction.Timestamp = util.CreateUtcTimestamp()

	return transaction, nil
}