package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/with-go/standard/array"
)

type Todo struct {
	id   int
	name string
	done bool
}

func main() {
	todoList := array.New()

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Add todo...")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Bye....")
			break
		}

		todo := Todo{
			id:   len(todoList) + 1,
			name: input,
			done: false,
		}

		if input == "todos" {
			fmt.Println(todoList)
		} else if input == "complete" {
			fmt.Println("Which task you want to complete")
			for {
				taskId, _ := reader.ReadString('\n')
				taskId = strings.TrimSpace(taskId)
				if taskId != "" {

					break
				}
			}
		} else {
			todoList = todoList.Push(todo)

		}
	}

}
