package models

type Product struct {
	 Title string `json:"id"`
	 Description string `json:"description"`
	 Image string `json:"image"`
	 Price string `json:"price"`
}

type ProductWrite struct {
	gorm.Model
	Product
}

// hides the gorm details from the API
type ProductRead struct {
	ID int `json:"id"`
	Product
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// place holder
func (u *Product) Validate() error {
	return nil
}