package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func InitLogger() {
	Log = log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)
}
