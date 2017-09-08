package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"github.com/hyperledger/fabric/core"
	"github.com/op/go-logging"
)

func checkDB() bool{
	if file,_ := ioutil.ReadFile("db.txt"); file == nil{
		fmt.Println("Не найден файл с данными!\nДля создания  файла с данными выполните приложение с аргументом <w> и файлом транзакции!")
		return false
	}else{return true}
}

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
}
func readTransaction(id string){
	if checkDB() {
		core.ReadTransaction(id)
	}
}
func testBlock(){
	if checkDB() {
		core.TestValidAllBlocks()
	}
}

func main(){
	core.Checksum()
	args := os.Args
	if len(args) < 2{
		fmt.Printf("Внимание, введите аргумент!")
	}else {
		if len(args) == 3 && strings.Compare(args[2],"debug") == 0{
			logging.SetLevel(logging.DEBUG, "")
		}else{
			logging.SetLevel(logging.NOTICE, "")
		}
		method := args[1]
		if strings.Compare(method, "i") == 0 {
			fmt.Printf("Начата инициализация базы данных.")
			createBlock()
		} else if strings.Compare(method, "w") == 0 {
			fmt.Printf("Начата запись блока.")
			readFile("./" + os.Args[2])
		} else if strings.Compare(method, "r") == 0 {
			fmt.Printf("Начато чтение транзакции.")
			readTransaction(os.Args[2])
		} else if strings.Compare(method, "t") == 0 {
			fmt.Printf("Начата проверка базы данных.")
			testBlock()
		} else {
			fmt.Printf("Внимание, неправильный аргумент!")
		}
	}
}