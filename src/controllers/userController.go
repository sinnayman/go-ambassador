package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"ambassador/src/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	db *database.Database
}

func NewUserController(db *database.Database) *UserController {
	return &UserController{db: db}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return utils.SendErrorResponse(ctx, "passwords don't match", fiber.StatusBadRequest)
	}

	user := models.UserWrite{
		FirstName:        data["first_name"],
		LastName:         data["last_name"],
		Email:            data["email"],
		PasswordValidate: data["password"],
		PasswordConfirm:  data["password_confirm"],
		IsAmbassador:     false,
	}

	// Validate the user
	if err := user.Validate(); err != nil {
		return utils.SendErrorResponse(ctx, err.Error(), fiber.StatusBadRequest)
	}

	if err := user.SetPassword(data["password"]); err != nil {
		return utils.SendErrorResponse(ctx, err.Error(), fiber.StatusBadRequest)
	}

	c.db.Create(&user)

	return ctx.JSON(user)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.UserRead
	result := c.db.DB.Where("email = ?", data["email"]).First(&user)

	if result.Error != nil {
		return utils.SendErrorResponse(ctx, "Invalid Credentials", fiber.StatusBadRequest)
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return utils.SendErrorResponse(ctx, "Invalid Credentials", fiber.StatusBadRequest)
	}

	expires := time.Now().Add(time.Hour * 24)
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.At(expires),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		return utils.SendErrorResponse(ctx, "Invalid Credentials", fiber.StatusBadRequest)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expires,
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return nil
}

func (c *UserController) Authenticate(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return utils.SendErrorResponse(ctx, "Invalid Authentication", fiber.StatusForbidden)
	}

	payload := token.Claims.(*jwt.StandardClaims)

	var user models.UserRead
	result := c.db.DB.Where("id = ?", payload.Subject).First(&user)
	if result.Error != nil {
		return utils.SendErrorResponse(ctx, result.Error.Error(), fiber.StatusBadRequest)
	}

	return ctx.JSON(user)
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	return nil
}
