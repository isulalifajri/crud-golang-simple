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

