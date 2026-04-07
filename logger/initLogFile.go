package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func initLogFile() *os.File {
	dir := "logs"
	now := time.Now()
	timeStamp := now.Format("2006-01-02 15:04:05.000")
	fileName := fmt.Sprintf("log-%s.log", timeStamp)
	filePath := filepath.Join(dir, fileName)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
