package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{"one", "two", "three"}

	ch := make(ExampleChannel)

	for i := range tasks {
		go ch.callChannel(tasks[i])
	}

	for i := 0; i < len(tasks); i++ {
		msg := <-ch
		fmt.Println(msg)
	}

	fmt.Println("wait")
}

type ExampleChannel chan string

func (ch ExampleChannel) callChannel(word string) {
	time.Sleep(time.Second * 4)
	ch <- word
}
