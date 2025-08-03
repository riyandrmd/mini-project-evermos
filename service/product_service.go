package service

import (
	"fmt"
	"strconv"
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

func GetAllProducts(filters map[string]string) ([]model.Product, error) {
	var products []model.Product
	db := config.DB.Preload("Toko").Preload("Category").Preload("Images")

	if q, ok := filters["nama_produk"]; ok && q != "" {
		db = db.Where("nama_produk LIKE ?", "%"+q+"%")
	}
	if cat, ok := filters["id_category"]; ok && cat != "" {
		db = db.Where("category_id = ?", cat)
	}
	if toko, ok := filters["id_toko"]; ok && toko != "" {
		db = db.Where("toko_id = ?", toko)
	}
	if min, ok := filters["harga_min"]; ok && min != "" {
		db = db.Where("harga >= ?", min)
	}
	if max, ok := filters["harga_max"]; ok && max != "" {
		db = db.Where("harga <= ?", max)
	}

	page, _ := strconv.Atoi(filters["page"])
	limit, _ := strconv.Atoi(filters["limit"])
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id string) (*model.Product, error) {
	var product model.Product

	err := config.DB.
		Preload("Toko").
		Preload("Category").
		Preload("Images").
		First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}
