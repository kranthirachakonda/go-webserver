package controllers

import (
	"net/http"
)

// RegisterControllers registers all the controllers and its handlers
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
