package main

import (
	"flag"
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

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "List all the books")

	fmt.Println(getAll)
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data is byte \n", string(dataByte))
}
