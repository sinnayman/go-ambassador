package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/database"
	"ambassador/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, db *database.Database) {
	api := app.Group("/api")

	admin := api.Group("/admin")

	userController := controllers.NewUserController(db)
	admin.Post("/register", userController.Register)
	admin.Get("/login", userController.Login)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("/User", userController.GetAuthenticatedUser)
	adminAuthenticated.Get("/logout", userController.Logout)
}
