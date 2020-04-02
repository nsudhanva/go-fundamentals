package controllers

import (
	"regexp"
	"net/http"
)

type userController struct{
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user controller!"))
}

// Constructor
// Go automatically promote the variable and not freeing memory

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile('^/users/(\d+)/?')
	}
}