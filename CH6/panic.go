package main

import "fmt"

// Program entered a state that cannot continue

func main() {
	fmt.Println("Starting web server...")

	panic("Something bad just happened")

	fmt.Println("Stopping web server...")
}
