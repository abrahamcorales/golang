package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	count int
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.count++
	fmt.Println("[LOG]:", message)
}

func main() {
	logger1 := GetLogger()
	logger2 := GetLogger()

	logger1.Log("First message")
	logger2.Log("Second message")

	fmt.Println("Logger instances are the same:", logger1 == logger2)
	fmt.Println("Log count:", logger1.count)
}
