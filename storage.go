package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func updateJsonFile() {
	formatList := make([]formatedTodo, 0)

	for i := range todoList {
		todo := todoList[i]
		formatList = append(formatList, formatedTodo{ID: todo.id, NAME: todo.name, DONE: todo.done})
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
