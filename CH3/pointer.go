package main

import (
	"fmt"
)

func main() {
	// var firstName *string
	var firstName *string = new(string)

	// nil

	// Error
	*firstName = "Sudhanva"
	fmt.Println(*firstName)

	lastName := "Sudhanva"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Narayana"
	fmt.Println(ptr, *ptr)
}
