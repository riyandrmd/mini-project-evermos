package model

import "gorm.io/gorm"

type LogProduct struct {
	gorm.Model
	ProductID     uint      `json:"product_id"`
	Product       *Product  `json:"-"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	PriceReseller float64   `json:"price_reseller"`
	PriceCustomer float64   `json:"price_customer"`
	CategoryID    uint      `json:"category_id"`
	Category      *Category `json:"-"`
	StoreID       uint      `json:"store_id"`
	Store         *Store    `json:"-"`
}
