package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"fmt"

	"github.com/go-faker/faker/v4"
)

func main() {

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to db")
		return
	}

	var count int64
	if err := db.DB.Model(&models.UserRead{}).Where("is_ambassador = ?", true).Count(&count).Error; err != nil {
		// Handle the error, e.g., log it or return an error response
		fmt.Println("Error counting records:", err)
		return
	}

	if count > 0 {
		fmt.Println("Data already exists, skipping seeding")
		return
	}

	for i := 0; i < 1; i++ {
		ambassador := models.UserWrite{
			User: models.User{
				FirstName:    faker.FirstName(),
				LastName:     faker.LastName(),
				Email:        faker.Email(),
				IsAmbassador: true,
			},
		}
		ambassador.SetPassword("1234")
		db.DB.Create(&ambassador)
	}
}
