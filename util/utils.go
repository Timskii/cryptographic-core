package util

import (
		"os"
		"strconv"
		//"time"
	)


func PrintData(data []byte){
	//dat := time.Now()
	filename := "db" //dat.Format("20060102")

	file, _ := os.OpenFile(filename+".txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}

func GenerateKey(value *[]string) string {
	var key string
	args := *value
	key = strconv.Itoa(len(args[0]))+args[0]+strconv.Itoa(len(args[1]))+args[1]
	return key
}