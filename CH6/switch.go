package main

import (
	"fmt"
)

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		fmt.Println("Get Request")

	case "POST":
		fmt.Println("POST Request")

	case "DELETE":
		fmt.Println("DELETE Request")
	case "PUT":
		fmt.Println("PUT Request")

	default:
		fmt.Println("Unhandled method")
	}
}
