package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"flag"
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-faker/faker/v4"
)

func main() {
	seedAmbassadors := flag.Bool("ambassadors", false, "Seed ambassadors")
	seedProducts := flag.Bool("products", false, "Seed products")
	flag.Parse()

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to db")
		return
	}

	if *seedAmbassadors {
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

	if *seedProducts {
		for i := 0; i < 1; i++ {
			product := models.ProductWrite{
				Product: models.Product{
					Title:       gofakeit.BuzzWord(),         // Generates a random buzzword as the product title
					Description: gofakeit.Sentence(5),        // Generates a random sentence as the product description
					Image:       gofakeit.ImageURL(100, 100), // Generates a random image URL
					Price:       rand.Float64(),
				},
			}
			db.DB.Create(&product)
		}
	}
}
