package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "passwords don't match"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName:        data["first_name"],
		LastName:         data["last_name"],
		Email:            data["email"],
		Password:         password,
		PasswordValidate: data["password"],
		PasswordConfirm:  data["password_confirm"],
		IsAmbassador:     false,
	}

	// Validate the user
	if err := user.Validate(); err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": err.Error()})
	}

	c.db.Create(&user)

	return ctx.JSON(user)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	result := c.db.DB.Where("email = ?", data["email"]).First(&user)

	if result.Error != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	expires := time.Now().Add(time.Hour * 24)
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.At(expires),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
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
