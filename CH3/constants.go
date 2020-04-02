package main

import (
	"fmt"
)

func main() {
	const pi = 3.1413
	fmt.Println(pi)

	// Implicit constant
	const x = 3
	const c int = 3
	fmt.Println(x + 1.2)
	fmt.Println(float32(c) + 1.2)

}
