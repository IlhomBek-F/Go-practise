
package main

import (
	"github.com/modood/table"
	"fmt"
)

type Users struct {
	Name string 
	Age int
	Status bool
}

func main() {
	users := []Users{{Name: "John", Age: 29, Status: true}}
    
	s := table.AsciiTable(users)

	fmt.Println(s)
}