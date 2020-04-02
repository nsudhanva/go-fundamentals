package main

import "fmt"

func main() {
	// Slice is not a fixed size entity

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// Resizing is taken care by go

	slice = append(slice, 4, 43, 5)
	fmt.Println(slice)

	// Slices of slices

	s2 := slice[1:]
	s3 := slice[:2]

	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)
}
