package storage

import (
	"GO/model"
	"encoding/json"
	"fmt"
	"os"
)

var TodoList = []model.Todo{}

func UpdateJsonFile() {
	formatList := make([]model.FormatedTodo, 0)

	for i := range TodoList {
		todo := TodoList[i]
		formatList = append(formatList, model.FormatedTodo{ID: todo.Id, NAME: todo.Name, DONE: todo.Done})
	}

	data, err := json.MarshalIndent(formatList, "", " ")

	if err != nil {
		fmt.Println("Error: While enconding todo to json")
		return
	}

	err = os.WriteFile("todos.json", data, 0644)

	if err != nil {
		fmt.Println("Error: While saving data to json")
	}
}
