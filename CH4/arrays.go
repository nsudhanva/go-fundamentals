package main

import "fmt"

func main() {

	// Arays ahve fixed sized
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])
	//  fmt.Println(arr[4]) - Out of bound error

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[1])
}
