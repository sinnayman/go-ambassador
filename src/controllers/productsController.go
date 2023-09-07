package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	db *database.Database
}

func NewProductController(db *database.Database) *ProductController {
	return &ProductController{db: db}
}

func (p *ProductController) GetAll(ctx *fiber.Ctx) error {
	var results []models.Product
	var products []models.ProductRead

	p.db.Find(&results)

	for _, product := range results {
		productRead := models.ProductRead{
			Product: product,
			// You can set CreatedAt, UpdatedAt, DeletedAt to appropriate values if needed
		}
		products = append(products, productRead)
	}

	return ctx.JSON(products)
}

func (p *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	product := models.ProductWrite{
		Product: models.Product{
			Description: data["description"],
			Image:       data["image"],
			Price:       data["price"],
		},
	}

	p.db.Create(&product)

	return ctx.JSON(product)
}

func (p *ProductController) GetById(ctx *fiber.Ctx) error {
	var product models.ProductRead

	idParam := ctx.Params("id")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	p.db.Where("id = ?", productID).Find(&product)

	return ctx.JSON(product)
}

func (p *ProductController) UpdateById(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	idParam := ctx.Params("id")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	product := models.ProductRead{
		ModelRead: models.ModelRead{
			ID: int(productID),
		},
		Product: models.Product{
			Description: data["description"],
			Image:       data["image"],
			Price:       data["price"],
		},
	}

	p.db.DB.Model(&product).Updates(&product)

	return ctx.JSON(product)
}

func (p *ProductController) DeleteById(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	idParam := ctx.Params("id")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	product := models.ProductRead{
		ModelRead: models.ModelRead{
			ID: int(productID),
		},
	}

	p.db.DB.Delete(&product)

	return nil
}
