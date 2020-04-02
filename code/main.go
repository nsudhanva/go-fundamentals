package main

import (
	"fmt"

	"github.com/nsudhanva/go-fundamentals/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Sudhanva",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
