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
	Height	uint64 			`json:"height"`
	Block 	protos.Block 	`json:"block"`
}

type DataJson struct {
	Key 	[]byte
	Value	[]byte
}

var db_blocks string = "db.txt"

func GetBlockNumberByTransaction(idx string) int{
	var blockNumber int = 0
	var blockDbs []DataJson
	data, err := ioutil.ReadFile(db_blocks)
	if err != nil {
		logger.Error("err = %v", err)
	}
	fileS := "[" + strings.TrimRight(string(data),",\n") + "]"
	fileS = strings.Replace(fileS,"}{","},{",-1)
	json.Unmarshal([]byte(fileS), &blockDbs)
	for _,data := range blockDbs{
		if strings.Contains(string(data.Key),idx){
			logger.Debugf("data.Key = %v",data.Key)
			break
		}
		if string(data.Key) == "blockCount"{
			blockNumber++
		}
	}
	return blockNumber-1
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