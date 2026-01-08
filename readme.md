## membuat program crud sederhana dengan golang

structur folder nya nanti:
```
golang-crud-simple/
│
├── config/
│   └── database.go
│
├── models/
│   └── user.go
│
├── handlers/
│   └── user_handler.go
│
├── seeders/
│   └── user_seeder.go
│
├── routes/
│   └── routes.go
│
├── main.go
└── go.mod

```

# langkah pertama

jalankan :
1. `mkdir golang-crud-simple` -> buat file dengan terminal
2. `cd golang-crud-simple` -> masuk kedalam filenya
3. `go mod init golang-crud-simple` -> init module

step 3 ini nanti menghasilkan file dengan nama `go.md` di dlm folder `golang-crud-simple`

file go pertama:
buat file `main.go`
```
package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")
}

```

jalankan: `go run main.go`
output : `Hello Golang`

## konsep dasar wajib paham

Konsep dasar (wajib paham)

Sedikit teori dulu (nggak bisa diskip):

🔹 package main

Entry point aplikasi

Kayak `public/index.php` di Laravel

🔹 func main()

Fungsi pertama yang dijalankan

🔹 import

Mirip `use / require`

## web server paling sederhana

Sekarang kita bikin server HTTP tanpa database dulu

ganti code `main.go` jadi seperti ini:

```
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Golang API")
	})

	// === Jalankan server ===
	fmt.Println("Server berjalan di http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}


```

karena kita pake mux, maka jalankan ini:
`go get github.com/gorilla/mux`

jalankan: `go run .`

akses di browser: `http://localhost:2000`
output: `Hello API Golang`

# buat model user

Buat model User: `models/user.go`
```
package main

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

kemudian ubah isi `main.go` jadi seperti ini:

```
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Budi", Email: "budi@mail.com"},
		{ID: 2, Name: "Siti", Email: "siti@mail.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("Server berjalan di http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}

```

## koneksikan dengan database

kita disini akan menggunkana `gorm` dan `postgreSQL`
karena itu kita jalankan perintah ini:
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

```

Config database pakai GORM:
`config/database.go`
```
package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=belajar_go port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	log.Println("Database connected (GORM)")
}

```

Model pakai GORM
`models/user.go`
```
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

```

coba seeder sederhana
`seeders/user_seeder.go`
```
package seeders

import (
	"log"

	"golang-crud-simple/config"
	"golang-crud-simple/models"
)

func SeedUsers() {
	users := []models.User{
		{Name: "Budi", Email: "budi@mail.com"},
		{Name: "Siti", Email: "siti@mail.com"},
	}

	for _, user := range users {
		config.DB.Create(&user)
	}

	log.Println("Seeder user selesai")
}

```

isi `routes/routes.go`
```
package routes

import (
	"github.com/gorilla/mux"
	"belajar-go/handlers"
)

func RegisterRoutes(router *mux.Router) {
	// User routes
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}

```

isi `handlers/user_handler.go`
```
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"belajar-go/config"
	"belajar-go/models"
)

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	config.DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

```
kemudian `main.go` jadi seperti ini:
```
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

```

buat database dulu :
```
psql -U postgres
CREATE DATABASE golang_crud_simple;

```

jalankan lagi aplikasi nya: `go run .`

# menggunakan swagger untuk test crud

jalankan: `go install github.com/swaggo/swag/cmd/swag@latest`

kemudian jalankan ini : `go get github.com/swaggo/http-swagger`

initialisasi swagger: `swag init`

kalau succes muncul folder:
```
docs/
 ├─ docs.go
 ├─ swagger.json
 └─ swagger.yaml

```

kemudian isi `main.go` jadi seperti ini:
```
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

```

di model `user.go` tambhakn ini:
```
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

type UserRequest struct {
	Name  string `json:"name" example:"Hening"`
	Email string `json:"email" example:"hening@mail.com"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


```

kemudian ubah `user_handler.go` jadi seperti ini:
```
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang-crud-simple/config"
	"golang-crud-simple/models"
)

//
// GET /users
//

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Find(&users)

	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//
// GET /users/{id}
//

// GetUser godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 404 {string} string
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
}

//
// POST /users
//

// CreateUser godoc
// @Summary Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User data"
// @Success 201 {object} models.UserResponse
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	json.NewDecoder(r.Body).Decode(&req)

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	config.DB.Create(&user)

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

//
// PUT /users/{id}
//

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserRequest true "User data"
// @Success 200 {object} models.UserResponse
// @Failure 404 {string} string
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var req models.UserRequest
	json.NewDecoder(r.Body).Decode(&req)

	user.Name = req.Name
	user.Email = req.Email

	config.DB.Save(&user)

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
}

//
// DELETE /users/{id}
//

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Failure 404 {string} string
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

```

jalankan lagi : `swag init`

jalankan: `go run .`
akses swagger: `http://localhost:2000/swagger/index.html`

