package main

import (
	"log"
	"os"
)

func ErrorLog() *log.Logger {
	return log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
}

func InfoLog() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}
