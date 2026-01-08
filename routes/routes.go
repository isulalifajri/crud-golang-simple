package routes

import (
	"github.com/gorilla/mux"
	"golang-crud-simple/handlers"
)

func RegisterRoutes(router *mux.Router) {
	// User routes
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
