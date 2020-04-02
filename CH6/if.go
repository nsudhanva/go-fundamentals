package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main(){
	u1 := User{
		ID: 1,
		FirstName: "Sudhanva",
		LastName: "Narayana",
	}

	u2 := User{
		ID: 2,
		FirstName: "Shreyas",
		LastName: "S",	
	}

	if u1 == u2{
		fmt.Println("Same user")
	}
	else if u1.FirstName == u2.FirstName {
		fmt.Println("Similar user")
	}
	else{
		fmt.Println("Not the same user")
	}
}
