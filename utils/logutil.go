package gpiutils

import (
	"log"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(msg string) {
	InfoLogger.Println(msg)
}

func Error(err error) {
	ErrorLogger.Println(err)
}
