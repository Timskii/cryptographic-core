package core

import (
		"fmt"
		"encoding/json"
		
		"github.com/cryptographic-core/util"

		)


func AddData (jsonobject []*Jsonobject){

	fmt.Printf("AddData \n")
	/*state := state.NewState()
	state.TxFinish("txID", true)*/
	fmt.Printf("Results: %v\n", jsonobject[0])
	fmt.Printf("function: %v\n", jsonobject[0].Params.CtorMsg.Function)


	data, _ := json.Marshal(jsonobject)
	util.PrintData(data)
}