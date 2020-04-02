package main

import (
	"net/http"

	"github.com/nsudhanva/go-fundamentals/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
