package main

import (

	"fmt"
	"os"


	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos"

)

func main() {
	transactions := []*protos.Transaction{}
	block := new(protos.Block)

	b := 1001
	 
	transaction, _ := protos.NewTransaction(	protos.ChaincodeID{Name: "chaincode_name"}, 
													"id_jfhdjsfhsjd" + string(b), 
													"test_function", 
													[]string{"one","two"})
	transactions = append(transactions,transaction)
	transactions = append(transactions,transaction)
	block.Transactions = transactions;
	mblock,_ := proto.Marshal(block)
	file, _ := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND , 0755)
	file.WriteString("\n--\n")
	file.WriteString(fmt.Sprintf("%#v\n", string(mblock)))
	file.WriteString("\n--\n")
	file.WriteString(fmt.Sprintf("%+v\n", *block))

	fmt.Printf("block v= %v \n v# = %#v \n v+ = %+v", block,block,block)
	fmt.Println("\n---------------------------------------------\n")
	fmt.Printf("mblock v= %v \n v# = %#v \n v+ = %+v", mblock,mblock,mblock)
	str_mblock := string(mblock)
	fmt.Println("\n---------------------------------------------\n")
	fmt.Printf("str_mblock v= %v \n v# = %#v \n v+ = %+v", str_mblock,str_mblock,str_mblock)

	block2 := new(protos.Block)
	err := proto.Unmarshal(mblock, block2)

	fmt.Println("\n---------------------------------------------\n")
	fmt.Println("err : %v", err)	
	fmt.Printf("block2 v= %v \n v# = %#v \n v+ = %+v", block2,block2,block2)



	
}