package controllers

import "net/http"

// RegisterController registers a new user controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}