package util

import (
		"os"
		"strconv"
		"encoding/json"

		//"github.com/hyperledger/fabric/protos"
		"github.com/hyperledger/fabric/protos"
		//"time"

	"io/ioutil"
	"strings"
)

type BlockDb struct {
	Height	uint64 			`json:"heigth"`
	Block 	protos.Block 	`json:"block"`

}
var filename string = "db_blocks.txt"

func GetBlocchainSize() uint64{
	var blockCount uint64 = 0
	var blockDbs []BlockDb

	dataCount, _ := ioutil.ReadFile(filename)
	fileS := "[" + strings.TrimRight(string(dataCount),",\n") + "]"
	json.Unmarshal([]byte(fileS), &blockDbs)
	blockCount = uint64(len(blockDbs))
	return blockCount
}


func PrintData(data []byte){
	//dat := time.Now()
	filename := "db"//"github.com/hyperledger/fabric/db" //dat.Format("20060102")

	file, _ := os.OpenFile(filename+".txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}

func PrintDataBlock(block protos.Block){
	var blockDb BlockDb

	blockDb.Height = GetBlocchainSize()
	blockDb.Block = block

	data,_ := json.Marshal(blockDb)
	data = append(data,[]byte(",\n")...)
	file, _ := os.OpenFile(filename,os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}

func GenerateKey(value *[]string) string {
	var key string
	args := *value
	key = strconv.Itoa(len(args[0]))+args[0]+strconv.Itoa(len(args[1]))+args[1]
	return key
}