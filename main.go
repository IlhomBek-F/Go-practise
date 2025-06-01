package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Id   int
	Name string
	Done bool
}

type formatedTodo struct {
	ID   int
	NAME string
	DONE bool
}

var todoList = []Todo{}

func main() {
	data, err := getTodoList()
	if err != nil {
		fmt.Println("Error: While getting todos from file")
	} else {
		for i := range data {
			todo := data[i]
			todoList = append(todoList, Todo{id: todo.ID, name: todo.NAME, done: todo.DONE})
		}
	}
	initTodo()
}

func initTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add new todo")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			exitTerminal()
			break
		}

		if input == "complete" {
			completeTodo()
		} else if input == "getTodoList" {
			data, err := getTodoList()
			if err != nil {
				fmt.Println("Error, while fetching data from file")
			} else {
				fmt.Println(data)
			}
		} else if input == "delete" {
			success, err := deleteTodo()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(success)
			}
		} else {
			addTodo(input)
		}

	}
}

func exitTerminal() {
	fmt.Println("Goodbye....")
}
