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

var Version string

//функция проверяет наличие файла с базой данных
func checkDB() bool{
	if file,_ := ioutil.ReadFile("db.txt"); file == nil{
		fmt.Println("Не найден файл с данными!\nДля создания  файла с данными выполните приложение с аргументом <w> и файлом транзакции!")
		return false
	}else{return true}
}

//функция по чтению файла во входящем параметра и записи его транзакции в базу данных
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

func main(){
	//Проверка несанкционированного изменения программы
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
			fmt.Printf("Начната инициализация базы даных.\n")
			core.CreateNilBlock()
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
				if checkDB() {
					core.ReadTransaction(args[2])
				}
			}
		} else if strings.Compare(method, "t") == 0 {
			fmt.Printf("Начата проверка базы данных.\n")
			if checkDB() {
				core.TestValidAllBlocks()
			}
		} else if strings.Compare(method,"h")==0 {
				fmt.Printf("Криптографическое ядро Hyperledger\n" +
									"\nИспользование:\n" +
									"\n	HLCK.exe [аргументы] <дополнительные аргументы>\n" +
									"\nДоступны аргументы:\n"+
									"\n	i - Инициализация базы данных\n"+
									"	w - Запись в блоки данных из файла\n"+
									"	r - Чтение транзакции для проверки\n"+
									"	t - Проверка базы данных\n" +
									"	v - Версия программы\n"+
									"\nДополнительные агрументы:\n" +
									"\n	debug - запустить приложение в режиме отладки\n\n")
		} else if strings.Compare(method, "v") == 0{
			fmt.Printf("Версия программы %v",Version)
		} else {
			fmt.Printf("Внимание, неправильный аргумент!\nДля получение справки запустите приложение с аргументом <h>.")
		}
	}
}