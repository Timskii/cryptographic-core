/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package protos

import (
	"encoding/json"
	"fmt"

	//"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/util"
)

// NewChaincodeExecute is used to invoke chaincode.
func NewChaincodeExecute(chaincodeInvocationSpec *ChaincodeInvocationSpec, uuid string, typ Transaction_Type) (*Transaction, error) {
	transaction := new(Transaction)
	transaction.Type = typ
	transaction.Txid = uuid
	transaction.Timestamp = util.CreateUtcTimestamp()
	cID := chaincodeInvocationSpec.ChaincodeSpec.GetChaincodeID()
	if cID != nil {
		data, err := json.Marshal(cID)
		if err != nil {
			return nil, fmt.Errorf("Could not marshal chaincode : %s", err)
		}
		transaction.ChaincodeID = data
	}
	data, err := json.Marshal(chaincodeInvocationSpec)
	if err != nil {
		return nil, fmt.Errorf("Could not marshal payload for chaincode invocation: %s", err)
	}
	transaction.Payload = data
	return transaction, nil
}

type strArgs struct {
	Function string
	Args     []string
}
