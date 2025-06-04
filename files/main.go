package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Example of content to file"

	file, err := os.Create("./example.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()

	readFile("example.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data is byte \n", string(dataByte))
}
