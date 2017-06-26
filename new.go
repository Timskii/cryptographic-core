package main
import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
	"strings"

	"github.com/tecbot/gorocksdb"
 
)





	
func main(){


	file, e := ioutil.ReadFile("db.txt")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}


	fileS := "[" + strings.TrimRight(string(file),",\n") + "]"
		
	
	var jsontype []*gorocksdb.DataJson
	json.Unmarshal([]byte(fileS), &jsontype)
	
	fmt.Println("---------------------------")
	fmt.Printf("jsontype %v\n", jsontype)
	fmt.Println("---------------------------")
	fmt.Printf("file %v\n", string(file))
	fmt.Printf("fileS %v\n", string(fileS))

}