package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

type AmbassadorController struct {
	db *database.Database
}

func NewAmbassadorController(db *database.Database) *AmbassadorController {
	return &AmbassadorController{db: db}
}

func (c *AmbassadorController) GetAll(ctx *fiber.Ctx) error {
	var users []models.UserRead
	c.db.DB.Where("is_ambassador = true").Find(&users)
	return ctx.JSON(users)
}
