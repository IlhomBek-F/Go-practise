package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	message string
}

type FileLogger struct {
	message string
}

func main() {
	consoleLogger := ConsoleLogger{message: "Console logger"}
	fileLogger := FileLogger{message: "File logger"}

	result, err := os.ReadFile("../todos.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s", string(result))
	processEvent(consoleLogger, consoleLogger.message)
	processEvent(fileLogger, fileLogger.message)
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (f FileLogger) Log(message string) {
	fmt.Print(message)
}

func processEvent(log Logger, message string) {
	log.Log(message)
}
