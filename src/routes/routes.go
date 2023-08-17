package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/database"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, db *database.Database) {
	api := app.Group("/api")

	admin := api.Group("/admin")

	userController := controllers.NewUserController(db)
	admin.Post("/register", userController.Register)
	admin.Get("/login", userController.Login)
}
