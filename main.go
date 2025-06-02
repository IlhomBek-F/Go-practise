package main

import (
	"GO/model"
	"GO/storage"
	"GO/todos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := todos.GetTodoList()
	if err != nil {
		fmt.Println("Error: While getting todos from file")
	} else {
		for i := range data {
			todo := data[i]
			storage.TodoList = append(storage.TodoList, model.Todo{Id: todo.ID, Name: todo.NAME, Done: todo.DONE})
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
			todos.CompleteTodo()
		} else if input == "getTodoList" {
			data, err := todos.GetTodoList()
			if err != nil {
				fmt.Println("Error, while fetching data from file")
			} else {
				fmt.Println(data)
			}
		} else if input == "delete" {
			success, err := todos.DeleteTodo()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(success)
			}
		} else {
			todos.AddTodo(input)
		}

	}
}

func exitTerminal() {
	fmt.Println("Goodbye....")
}
