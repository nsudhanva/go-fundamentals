package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for k, v := range wellKnownPorts {
		// Mandatroy to use the variable here
		fmt.Println(k, v)
	}
}
