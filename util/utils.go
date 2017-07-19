package util

import (
		"os"
		"strconv"
		"encoding/json"

		//"github.com/hyperledger/fabric/protos"
		"github.com/hyperledger/fabric/protos"
		//"time"
	)

type BlockDb struct {
	Height	int32 			`json:"heigth"`
	Block 	protos.Block 	`json:"block"`

}


func PrintData(data []byte){
	//dat := time.Now()
	filename := "db"//"github.com/hyperledger/fabric/db" //dat.Format("20060102")

	file, _ := os.OpenFile(filename+".txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}

func PrintDataBlock(block protos.Block){
	//dat := time.Now()
	filename :="db_blocks"  //"github.com/hyperledger/fabric/db_blocks" //dat.Format("20060102")
	var blockDb BlockDb
	blockDb.Height = 1
	blockDb.Block = block

	data,_ := json.Marshal(blockDb)
	data = append(data,[]byte(",\n")...)

	file, _ := os.OpenFile(filename+".txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}

func GenerateKey(value *[]string) string {
	var key string
	args := *value
	key = strconv.Itoa(len(args[0]))+args[0]+strconv.Itoa(len(args[1]))+args[1]
	return key
}