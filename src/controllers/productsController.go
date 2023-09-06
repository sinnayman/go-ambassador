package controllers

type ProductController struct {
	db *database.Database
}

func NewProductController(db *database.Database) *ProductController {
	return &ProductController{db: db}
}

func (p *ProductController) GetAll(ctx *fiber.Ctx)  error {
	 var products []models.Product

	 c.db.Find(&products)

	 return ctx.JSON(products)
}

func (p *productController) CreateProduct(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	product := models.ProductWrite{
		User: models.Product{
			Description:        data["description"],
			Image:         		data["image"],
			Price:            	data["price"],
		},
	}

	c.db.Create(&product)

	return ctx.JSON(product)
}

func (p *ProductController) GetById(ctx *fiber.Ctx)  error {
	var product models.ProductRead

    idParam := ctx.Params("id")
    productID, err := strconv.Atoi(idParam)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",``
        })
    }

	c.db.Where("id = ?", productID).Find(&product)

	return ctx.JSON(product)
}

func (p *ProductController) UpdateById(ctx *fiber.Ctx)  error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	idParam := ctx.Params("id")
    productID, err := strconv.Atoi(idParam)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",``
        })
    }

	product := models.ProductRead{
		ID: int(productID),
		Product: models.Product{
			Description: 	data["description"],
			Image:  		data["image"],
			Price:     		data["price"],
		},
	}

	c.db.DB.Model(&product).Updates(&product)

	return ctx.JSON(product)
}

func (p *ProductController) DeleteById(ctx *fiber.Ctx)  error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	idParam := ctx.Params("id")
    productID, err := strconv.Atoi(idParam)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",``
        })
    }

	product := models.ProductRead{
		ID: int(productID)
	}

	c.db.DB.Delete(&product)

	return nil
}