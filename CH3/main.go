package main

import (
	"fmt"
)

func main() {
	// Primitive decleration
	var i int
	i = 42
	fmt.Println(i)

	// Implicit decleration
	var f float32 = 3.14
	fmt.Println(f)
	// fmt.Println(f)
	// Complie time error - variable declared but not used just like packages

	// Explicit decleration - commonly used
	firstName := "Sudhanva"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	// Complex
	c := complex(3, 4)
	fmt.Println(c)

	// Multiple assignments
	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
