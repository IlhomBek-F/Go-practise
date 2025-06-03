package main

import (
	"GO/model"
	"GO/storage"
	"GO/todos"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	Name   string
	Age    int
	Status bool
}

func main() {
	db := connectToDB()
	var id int

	query := `INSERT INTO todos (name, completed) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, "test item", false).Scan(&id)

	if err != nil {
		log.Fatal("Error inserting todo into database:", err)
	}

	fmt.Println("✔ Todo inserted successfully with ID:", id)

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

func connectToDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return db
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
