package helper

import (
	"fmt"
	"models/model"
)

func findIndex(list []Todo, id int) int {
	for i := range list {
		fmt.Println(i)
		if list[i].ID == id {
			return i
		}
	}
	return -1
}

func filter(cb func(model.Todo) bool) []model.Todo {
	filteredList := []model.Todo{}
	for i := range todoList {
		if cb(todoList[i]) {
			filteredList = append(filteredList, todoList[i])
		}
	}

	return filteredList
}
