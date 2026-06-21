package models

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       uint           `json:"price"`

	Images []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}