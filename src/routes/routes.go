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
	ambassadorController := controllers.NewAmbassadorController(db)
	admin.Post("/", userController.Register)
	admin.Get("/login", userController.Login)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("/", userController.GetAuthenticatedUser)
	adminAuthenticated.Get("/logout", userController.Logout)
	adminAuthenticated.Put("/", userController.UpdateUser)
	adminAuthenticated.Put("/password", userController.UpdatePassword)

	adminAuthenticated.Get("/ambassadors", ambassadorController.GetAll)

	adminAuthenticated.Get("/products", productsController.GetAll)
	adminAuthenticated.Post("/products", productsController.CreateProduct)
	adminAuthenticated.Get("/products/:id", productsController.GetById)
	adminAuthenticated.Put("/products/:id", productsController.UpdateById)
	adminAuthenticated.Put("/products/:id", productsController.DeleteById)
}
