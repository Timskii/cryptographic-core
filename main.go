package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"github.com/hyperledger/fabric/core"

	"time"
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
	if file,_ := ioutil.ReadFile("db.txt"); file == nil{
		fmt.Println("Не найден файл с данными!\nДля создания  файла с данными выполните приложение с аргументом <w> и файлом транзакции!")
	}else {
		core.ReadTransaction(id)
	}
}
func testBlock(){
	fmt.Println("Test begin ", time.Now())
	core.TestValidAllBlocks()
	fmt.Println("Test end ",time.Now())
}

func main(){
	fmt.Printf("Args %v\n", os.Args)
	method := os.Args[1]
	if strings.Compare(method,"i")==0 {
		createBlock()
	} else if strings.Compare(method,"w") ==0 {
		readFile("./"+os.Args[2])
	} else if strings.Compare(method,"r")==0 {
		readTransaction(os.Args[2])
	} else if strings.Compare(method,"t")==0 {
		testBlock()
	} else if strings.Compare(method,"checksum")==0{
		core.Checksum()
	}
}