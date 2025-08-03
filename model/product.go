package model

import "time"

type Product struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	NamaProduk string `json:"nama_produk"`
	Slug       string `json:"slug"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`

	CategoryID uint `json:"id_category"`
	TokoID     uint `json:"id_toko"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Toko     Toko           `json:"toko" gorm:"foreignKey:TokoID"`
	Category Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Images   []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}
