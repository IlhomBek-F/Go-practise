package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserInfo struct {
	Name     string
	LastName string
	Age      int
}

func main() {
	callUserINfo()
	sliceArray()

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}

func sliceArray() {
	primes := [4]int{3, 4, 5, 6}

	var s []int = primes[0:2]

	fmt.Println(s)
}

func callUserINfo() {
	users := []UserInfo{{Name: "John", LastName: "Doe", Age: 30}, {Name: "Jane", LastName: "Smith", Age: 25}}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Name)
	}
	fmt.Println(("User Information:"))
}
