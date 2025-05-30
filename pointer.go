package main

import "fmt"

type Pointer struct {
	name string
}

func pointer() {
	var b = 5
	var a *int
	a = &b
	*a = 90
	fmt.Println(*a, b)
}
