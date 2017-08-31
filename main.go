package main
import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
	"strings"


	"github.com/hyperledger/fabric/core"
 
)


func readFile (fileName string){
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

func createBlock(){
	core.CreateNilBlock()
	fmt.Println("Block is created")
}

func readTransaction(id string){
	core.ReadTransaction(id)
	fmt.Printf("transaction id = %v\n", id)
}

func testBlock(){
	fmt.Println("Test begin")

	fmt.Println("Test end")

}

	
func main(){
	fmt.Printf("Args1 %v\n", os.Args)
	method := os.Args[1]

	if strings.Compare(method,"i")==0 {
		createBlock()
	} else if strings.Compare(method,"w") ==0 {
		readFile("./"+os.Args[2])
	} else if strings.Compare(method,"r")==0 {	
		readTransaction(os.Args[2])
	} else if strings.Compare(method,"t")==0 {
		testBlock()
	}
}