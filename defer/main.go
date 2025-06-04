package main

import (
	"fmt"
)

func main() {
	deferInLoop()
}

func deferInLoop() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Hello
// Two
// One
// WOrld
