package models

func (ProductWrite) TableName() string {
	return "products"
}

type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type ProductWrite struct {
	ModelWrite
	Product
}

// hides the gorm details from the API
type ProductRead struct {
	ModelRead
	Product
}

// place holder
func (u *Product) Validate() error {
	return nil
}
