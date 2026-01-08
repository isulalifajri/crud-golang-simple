package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang-crud-simple/config"
	"golang-crud-simple/models"
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
