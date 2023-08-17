package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// cors enables a browser to make a request to an endpoint port different than the web one (80)
	// the enable credentials option allows to also pass the cookie to that request
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	db, _ := database.NewDatabase()
	db.AutoMigrate()
	defer db.Close()

	routes.Setup(app, db)

	app.Listen(":3000")
}
