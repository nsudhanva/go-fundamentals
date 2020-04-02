package main

import (
	"fmt"
)

// Four types of looping
// No while or do while loops

func main() {
	var i int

	for i < 5 {
		fmt.Println(i)
		i++

		if i == 3 {
			continue
		}

		if i == 3 {
			break
		}

	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	// for ; ; {
	// 	fmt.Println("Infinite")
	// }
}
