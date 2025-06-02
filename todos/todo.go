package todos

import (
	"GO/helper"
	"GO/model"
	"GO/storage"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTodoList() ([]model.FormatedTodo, error) {
	data, err := os.ReadFile("todos.json")

	if err != nil {
		return nil, err
	}

	var todos = []model.FormatedTodo{}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func AddTodo(name string) {
	todo := model.Todo{
		Id:   len(storage.TodoList) + 1,
		Name: name,
		Done: false,
	}
	storage.TodoList = append(storage.TodoList, todo)
	storage.UpdateJsonFile()
}

func DeleteTodo() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Which task you want to delete %v? Enter task id", storage.TodoList)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid type id")
	}

	foundTaskId := helper.FindIndex(storage.TodoList, taskId)

	if foundTaskId != -1 {
		storage.TodoList = helper.Filter(func(todo model.Todo) bool {
			return todo.Id != storage.TodoList[foundTaskId].Id
		}, storage.TodoList)
		storage.UpdateJsonFile()
	} else {
		return "", errors.New("not todo found by id")
	}

	return "Todo deleted successfully", nil
}

func CompleteTodo() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Which task you want to complete %v? Enter task id", storage.TodoList)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		taskId, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid type id")
		}

		foundTodoIndex := helper.FindIndex(storage.TodoList, taskId)

		if foundTodoIndex != -1 {
			completedTodo := storage.TodoList[foundTodoIndex]
			completedTodo.Done = true

			storage.TodoList[foundTodoIndex] = completedTodo

			storage.UpdateJsonFile()
			break
		} else {
			fmt.Printf("Task not found with %v", taskId)
		}
	}

	fmt.Println("Success, Task completed")
}
