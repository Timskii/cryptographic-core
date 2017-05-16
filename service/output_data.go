package service

import (
		"fmt"
		"os"
		"time"
	)

func printData(data []byte){
	dat := time.Now()
	filename := dat.Format("20060102")

	file, _ := os.OpenFile(filename+".txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.close()
	_, _ = file.Write(data)
}