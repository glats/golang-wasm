package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *log.Logger
)

// GetLogger create new logger
func GetLogger() *log.Logger {
	once.Do(func() {
		logger = log.New(os.Stdout, "http: ", log.LstdFlags)
	})
	return logger
}
