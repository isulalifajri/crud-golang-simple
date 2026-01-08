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
