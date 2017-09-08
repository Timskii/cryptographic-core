package main
import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"

	"github.com/hyperledger/fabric/core"
	"github.com/op/go-logging"

)


func main(){
	logging.SetLevel(logging.DEBUG, "")
	//readFileTransfer()
	//core.CreateNilBlock()
	//readFileRegister()
	testBlocks()
	//core.Checksum()
}

func readFileTransfer(){
	filename := "github.com/hyperledger/fabric/transactions_transfer1.json"
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsontype []*core.Jsonobject
	json.Unmarshal(file, &jsontype)
	core.AddData(jsontype)
}

func readFileRegister(){
	filename := "github.com/hyperledger/fabric/transactions_register.json"
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsontype []*core.Jsonobject
	json.Unmarshal(file, &jsontype)
	core.AddData(jsontype)
}
func testBlocks()  {
	core.TestValidAllBlocks()
}