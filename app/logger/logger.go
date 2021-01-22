package logger

import (
	"io"
	"io/ioutil"
	"log"
)

var (
	defaultlogger = NewLogger(ioutil.Discard)
)

type logger struct {
	logger *log.Logger
}

// NewLogger :
func NewLogger(w io.Writer) *logger {

	l := log.New(w, "default", log.Ldate|log.Ltime|log.Lshortfile)
	return &logger{
		logger: l,
	}
}

// Info ...
func (l *logger) Info(content string) {
	l.logger.Println("Info:" + content)
}

// Default Info ...
func Info(content string) {
	defaultlogger.Info(content)
}

// Warn ...
func (l *logger) Warn(content string) {
	l.logger.Println("Warn:" + content)
}

// Default Info ...
func Warn(content string) {
	defaultlogger.Warn(content)
}

// Error ...
func (l *logger) Error(content string) {
	l.logger.Println("Error:" + content)
}

// Default Info ...
func Error(content string) {
	defaultlogger.Error(content)
}

func SetDefaultLogger(l *logger) {
	defaultlogger = l
}
