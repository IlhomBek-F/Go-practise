package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getTodoList() ([]formatedTodo, error) {
	data, err := os.ReadFile("todos.json")

	if err != nil {
		return nil, err
	}

	var todos = []formatedTodo{}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func addTodo(name string) {
	todo := Todo{
		id:   len(todoList) + 1,
		name: name,
		done: false,
	}
	todoList = append(todoList, todo)
	updateJsonFile()
}

func deleteTodo() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Which task you want to delete %v? Enter task id", todoList)

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
		updateJsonFile()
	} else {
		return "", errors.New("not todo found by id")
	}

	return "Todo deleted successfully", nil
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
			updateJsonFile()
			break
		} else {
			fmt.Printf("Task not found with %v", taskId)
		}
	}

	fmt.Println("Success, Task completed")
}
