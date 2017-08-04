package main
import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"



	"github.com/hyperledger/fabric/core"
 
)


func readFile1 (fileName string){
	fmt.Printf("readFile start\n")

	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsontype []*core.Jsonobject
	json.Unmarshal(file, &jsontype)
	core.AddData(jsontype)
}



	
func main(){
	fmt.Printf("Args1 %v\n", os.Args)

	file, e := ioutil.ReadFile("github.com/hyperledger/fabric/transactions.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsontype []*core.Jsonobject
	json.Unmarshal(file, &jsontype)
	core.AddData(jsontype)


}