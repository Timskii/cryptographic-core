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
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("Внимание, при чтении файла возникли ошибки: %v\n", e)
		os.Exit(1)
	}
	var jsontype []*core.Jsonobject
	json.Unmarshal(file, &jsontype)
	if len(jsontype)>0 {
		core.AddData(jsontype)
	}else{
		fmt.Printf("Внимание, неудалось прочитать транзакции!")
	}
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
		fmt.Printf("Внимание, введите аргумент!\nДля получение справки запустите приложение с аргументом <h>.")
	}else {
		for _, str := range args{
			if strings.Compare(str,"debug")==0 {
				logging.SetLevel(logging.DEBUG, "")
				break
			}else{logging.SetLevel(logging.NOTICE, "")}
		}

		method := args[1]
		if strings.Compare(method, "i") == 0 {
			fmt.Printf("Начата инициализация базы данных.\n")
			createBlock()
		} else if strings.Compare(method, "w") == 0 {
			if len(args) < 3{
				fmt.Printf("Внимание, укажите файл с транзакциями!")
			}else {
				fmt.Printf("Начата запись блока.\n")
				readFile("./" + args[2])
			}
		} else if strings.Compare(method, "r") == 0 {
			if len(args) < 3{
				fmt.Printf("Внимание, укажите ID транзакции!")
			}else {
				fmt.Printf("Начато чтение транзакции.\n")
				readTransaction(args[2])
			}
		} else if strings.Compare(method, "t") == 0 {
			fmt.Printf("Начата проверка базы данных.\n")
			testBlock()
		} else if strings.Compare(method,"h")==0 {
				fmt.Printf("Криптографическое ядро Hyperledger\n" +
									"\nИспользование:\n" +
									"\n	main.exe [аргументы]<дополнительные аргументы>\n" +
									"\nДоступны аргументы:\n"+
									"\n	i - Инициализация базы данных\n"+
									"	w - Запись в блоки данных из файла\n"+
									"	r - Чтение транзакции для проверки\n"+
									"	t - Проверка базы данных\n"+
									"\nДополнительные агрументы:\n" +
									"\n	debug - запустить приложение в режиме отладки\n\n")
		} else {
			fmt.Printf("Внимание, неправильный аргумент!\nДля получение справки запустите приложение с аргументом <h>.")
		}
	}
}