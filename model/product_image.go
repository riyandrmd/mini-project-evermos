package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID uint     `json:"product_id"`
	ImageURL  string   `json:"image_url"`
	Product   *Product `json:"-" gorm:"foreignKey:ProductID"`
}
