package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	port, err := startWebServer(port, 2)
	fmt.Println(port, err)

	//  startWebServer function as an object
}

// No neet to pass datatype in function paramter is both are of the same datatype - one datatype is required
func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server on port", port)
	fmt.Println("Number of retires", numberOfRetries)
	fmt.Println("Stopping server...")
	// return errors.New("Something went wrong")
	return port, nil
}
