package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	PriceReseller float64   `json:"price_reseller"`
	PriceCustomer float64   `json:"price_customer"`
	Description   string    `json:"description"`
	Stock         int       `json:"stock"`
	StoreID       uint      `json:"store_id"`
	CategoryID    uint      `json:"category_id"`
	Store         *Store    `json:"-"`
	Category      *Category `json:"-"`
}
