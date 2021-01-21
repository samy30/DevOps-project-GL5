package logger

import (
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info ...
func Info(content string) {
	infoLogger.Println(content)
}

// Warn ...
func Warn(content string) {
	warningLogger.Println(content)
}

// Error ...
func Error(content string) {
	errorLogger.Println(content)
}
