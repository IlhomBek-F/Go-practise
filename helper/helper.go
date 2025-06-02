package helper

import (
	"GO/model"
	"fmt"
)

func FindIndex(list []model.Todo, id int) int {
	for i := range list {
		fmt.Println(i)
		if list[i].Id == id {
			return i
		}
	}
	return -1
}

func Filter(cb func(model.Todo) bool, todoList []model.Todo) []model.Todo {
	filteredList := []model.Todo{}
	for i := range todoList {
		if cb(todoList[i]) {
			filteredList = append(filteredList, todoList[i])
		}
	}

	return filteredList
}
