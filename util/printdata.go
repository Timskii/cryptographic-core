package util

import (
		"os"
		"time"
	)


func PrintData(data []byte){
	dat := time.Now()
	filename := dat.Format("20060102")

	file, _ := os.OpenFile(filename+".txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.Write(data)
}