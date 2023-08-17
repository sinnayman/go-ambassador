package utils

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(ctx *fiber.Ctx, message string, statusCode int) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"message": message,
	})
}
