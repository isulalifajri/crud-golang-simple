// @title Golang CRUD Simple API
// @version 1.0
// @description CRUD API using Golang, GORM, PostgreSQL
// @host localhost:2000
// @BasePath /

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang-crud-simple/docs"

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
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	routes.RegisterRoutes(router)

	fmt.Println("Server berjalan di http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}
