package main

import "fmt"

func main() {
	// Zero is initialized for int, blank for string

	// Struct is in the scope of main function
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Sudhanva"
	u.LastName = "Narayana"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1, 
		FirstName: "Sudhanva", 
		LastName: "Dent"
	}

	fmt.Println(u2)
}
