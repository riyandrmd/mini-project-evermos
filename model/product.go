package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	StoreID    uint      `json:"store_id"`
	Store      *Store    `json:"-"`
	CategoryID uint      `json:"category_id"`
	Category   *Category `json:"-"`
}
