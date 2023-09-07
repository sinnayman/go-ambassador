package models

type Product struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       string `json:"price"`
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
