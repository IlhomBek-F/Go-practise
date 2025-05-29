package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	id   int
	name string
	done bool
}

var todoList = []Todo{}

func main() {
	// initTodo();
	userId := Todo{id: 1, name: "John", done: false}

	fmt.Println(userId.getTodoList())
}

func (e Todo) getTodoList() int {
	return e.id
}

func findIndex(list []Todo, id int) int {
	for i := range list {
		fmt.Println(i)
		if list[i].id == id {
			return i
		}
	}
	return -1
}

func filter(cb func(Todo) bool) []Todo {
	filteredList := []Todo{}
	for i := range todoList {
		if cb(todoList[i]) {
			filteredList = append(filteredList, todoList[i])
		}
	}

	return filteredList
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
			fmt.Println(todoList)
		} else if input == "delete" {
			deleteTodo()
		} else {
			addTodo(input)
		}

	}
}

func addTodo(name string) {
	todo := Todo{
		id:   len(todoList) + 1,
		name: name,
		done: false,
	}

	todoList = append(todoList, todo)
}

func deleteTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Which task you want to delete %v? Enter task id", todoList)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		taskId, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid type id")
		}

		foundTaskId := findIndex(todoList, taskId)

		if foundTaskId != -1 {
			todoList = filter(func(todo Todo) bool {
				return todo.id != todoList[foundTaskId].id
			})
			break
		}
	}

	fmt.Println(todoList)
}

func completeTodo() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Which task you want to complete %v? Enter task id", todoList)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		taskId, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid type id")
		}

		foundTodoIndex := findIndex(todoList, taskId)

		if foundTodoIndex != -1 {
			completedTodo := todoList[foundTodoIndex]
			completedTodo.done = true

			todoList[foundTodoIndex] = completedTodo
			break
		} else {
			fmt.Printf("Task not found with %v", taskId)
		}
	}

	fmt.Println("Success, Task completed")
}

func exitTerminal() {
	fmt.Println("Goodbye....")
}
