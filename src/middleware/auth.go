package middleware

import (
	"ambassador/src/utils"
	"strconv"

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

func GetUserID(ctx *fiber.Ctx) uint64 {
	uidStr, ok := ctx.Locals("User_ID").(string)
	if !ok {
		return 0
	}
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return 0
	}
	return uid
}
