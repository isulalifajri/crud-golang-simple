package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang-crud-simple/config"
	"golang-crud-simple/models"
	"golang-crud-simple/routes"
	"golang-crud-simple/seeders"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
	seeders.SeedUsers()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("Server berjalan di http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}
