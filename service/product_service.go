package service

import (
	"fmt"
	"strings"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func SaveProductImage(productID uint, url string) error {
	img := model.ProductImage{
		ProductID: productID,
		URL:       url,
	}
	return config.DB.Create(&img).Error
}

func CreateProduct(userID uint, input dto.CreateProductRequest) (*model.Product, error) {
	var toko model.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return nil, fmt.Errorf("toko tidak ditemukan")
	}

	slug := strings.ToLower(strings.ReplaceAll(input.NamaProduk, " ", "-"))

	product := model.Product{
		NamaProduk: input.NamaProduk,
		Slug:       slug,
		Deskripsi:  input.Deskripsi,
		Harga:      input.Harga,
		Stok:       input.Stok,
		TokoID:     toko.ID,
		CategoryID: input.CategoryID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
