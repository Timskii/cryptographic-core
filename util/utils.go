package util

import (
		"os"
		"strconv"
		"encoding/json"
		"io/ioutil"
		"strings"

		"github.com/hyperledger/fabric/protos"
		"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("utils.go")

type BlockDb struct {
	Height	uint64 			`json:"heigth"`
	Block 	protos.Block 	`json:"block"`

}
var db_blocks string = "db_blocks.txt"
/*
func GetBlockhainSize() uint64{
	var blockCount uint64 = 0
	var blockDbs []BlockDb

	dataCount, _ := ioutil.ReadFile(db_blocks)
	if dataCount != nil {
		fileS := "[" + strings.TrimRight(string(dataCount), ",\n") + "]"
		json.Unmarshal([]byte(fileS), &blockDbs)
		blockCount = uint64(len(blockDbs))
	}
	return blockCount
}*/

func GetBlockNumberByTransaction(idx string) int{
	var blockNumber int = 0
	var blockDbs []BlockDb
	data, _ := ioutil.ReadFile(db_blocks)
	fileS := "[" + strings.TrimRight(string(data),",\n") + "]"
	json.Unmarshal([]byte(fileS), &blockDbs)
	for i,block := range blockDbs{
		for _,transaction := range block.Block.Transactions {
			if  strings.Compare(transaction.Txid,idx) ==0 {
				blockNumber = i
				break
			}
		}
	}
	return blockNumber
}


func PrintData(data []byte,filename string){
	//dat := time.Now()
	file, _ := os.OpenFile(filename+".txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()
	_, _ = file.Write(data)
}


func PrintDataBlock(block protos.Block, size uint64){
	var blockDb BlockDb

	blockDb.Height = size
	blockDb.Block = block
	data,_ := json.Marshal(blockDb)
	logger.Debugf("%s\n",string(data))
}

func GenerateKey(value1, value2 string) string {
	var key string
	key = strconv.Itoa(len(value1))+value1+strconv.Itoa(len(value2))+value2
	return key
}

func ToChaincodeArgs(args []string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}