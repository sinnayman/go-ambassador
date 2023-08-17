package middleware

import (
	"ambassador/src/utils"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return utils.SendErrorResponse(ctx, "Invalid Authentication", fiber.StatusForbidden)
	}

	payload := token.Claims.(*jwt.StandardClaims)
	ctx.Locals("User_ID", payload.Subject)

	return ctx.Next()
}
