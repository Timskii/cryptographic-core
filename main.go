package main
import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
   // "github.com/hyperledger/fabric/core/ledger/statemgmt/state"
   "strings"
   "time"
)

type Jsonobject struct {
    Jsonrpc	string		`json:"jsonrpc"`
	Method	string		`json:"method"`
	Params	ParamsType	`json:"params"`
	Id	string			`json:"id"`
}

type ParamsType struct{
	Type int						`json:"type"`
	ChaincodeID 	ChaincodeIDType	`json:"chaincodeID"`
	CtorMsg			CtorMsgType		`json:"ctorMsg"`
}

type ChaincodeIDType struct{
	Name string			`json:"name"`
	}

type CtorMsgType struct {
	Function string		`json:"function"`
	Args []string		`json:"args"`
}

func createBlock(){
	fmt.Println("Block is created")
}

func readFile (fileName string){
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsontype []Jsonobject
	json.Unmarshal(file, &jsontype)
	/*state := state.NewState()

	state.TxFinish("txID", true)*/

	fmt.Printf("Results: %v\n", jsontype)
	fmt.Printf("function: %v\n", jsontype[0].Params.CtorMsg.Function)
	fmt.Printf("Args: %v\n", jsontype[0].Params.CtorMsg.Args[1])
}

func readTransaction(id string){
	fmt.Printf("transaction id = %v\n", id)
}

func testBlock(){
	fmt.Println("Test begin")
	fmt.Println("time ", time.Now())
	fmt.Println("weekday ", time.Now().Weekday())

	s := make([]string,3)
	s[1] = "ee"
	s = append(s, "ww")
	fmt.Println("s1 ",s , "len ", len(s))
	d := append(s,"efe", "efff")
	
	fmt.Println("d ",d , "len ", len(d))
	d = append(s)
	fmt.Println("s2 ",s , "len ", len(s))
	
	fmt.Println("d2 ",d , "len ", len(d))
}

	
func main(){
	fmt.Printf("Args1 %v\n", os.Args)
	method := os.Args[1]

	if strings.Compare(method,"c")==0 {
		createBlock()
	} else if strings.Compare(method,"w") ==0 {
		readFile("./"+os.Args[2])
	} else if strings.Compare(method,"r")==0 {	
		readTransaction(os.Args[2])
	} else if strings.Compare(method,"t")==0 {
		testBlock()
	}
	
	/*fmt.Printf("Compare r %v\n",strings.Compare(os.Args[1],"r"))
	fmt.Printf("Compare d %v\n",strings.Compare(os.Args[1],"d"))*/
}