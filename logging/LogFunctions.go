package logging

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime)
)

func Info(v ...any) {
	infoLogger.Println(v)
}

func Error(v ...any) {
	errorLogger.Println(v)
}

func Fatal(v ...any) {
	errorLogger.Fatalln(v)
}
