package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once         sync.Once
	httpLogger   *log.Logger
	systemLogger *log.Logger
)

// GetHttpLogger create logger for http requests
func GetHttpLogger() *log.Logger {
	once.Do(func() {
		httpLogger = log.New(os.Stdout, "http: ", log.LstdFlags)
	})
	return httpLogger
}

// GetSystemLogger create logger for system errors
func GetSystemLogger() *log.Logger {
	once.Do(func() {
		systemLogger = log.New(os.Stdout, "system: ", log.LstdFlags)
	})
	return systemLogger
}
