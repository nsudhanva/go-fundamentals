package main

import (
	"fmt"
)

// iota helps us to evolve a constant through constant block and constant expressions

const (
	first = iota + 6
	second = 2 << iota
)

const (
	third = iota
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)
}
