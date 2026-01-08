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
