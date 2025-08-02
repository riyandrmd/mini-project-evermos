package service

import (
	"errors"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateProduct(userID uint, request dto.CreateProductRequest) error {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return errors.New("store not found")
	}

	var exists model.Product
	if err := config.DB.Where("slug = ?", request.Slug).First(&exists).Error; err == nil {
		return errors.New("slug already exists")
	}

	product := model.Product{
		Name:          request.Name,
		Slug:          request.Slug,
		PriceReseller: request.PriceReseller,
		PriceCustomer: request.PriceCustomer,
		Description:   request.Description,
		CategoryID:    request.CategoryID,
		StoreID:       store.ID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return errors.New("failed to create product")
	}

	return nil
}

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(userID uint, id string, request dto.UpdateProductRequest) error {
	var product model.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	var store model.Store
	if err := config.DB.Where("id = ? AND user_id = ?", product.StoreID, userID).First(&store).Error; err != nil {
		return errors.New("unauthorized")
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Slug != "" {
		product.Slug = request.Slug
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.PriceReseller != 0 {
		product.PriceReseller = request.PriceReseller
	}
	if request.PriceCustomer != 0 {
		product.PriceCustomer = request.PriceCustomer
	}
	if request.CategoryID != 0 {
		product.CategoryID = request.CategoryID
	}

	if err := config.DB.Save(&product).Error; err != nil {
		return errors.New("failed to update product")
	}

	return nil
}

func DeleteProduct(userID uint, id string) error {
	var product model.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	var store model.Store
	if err := config.DB.Where("id = ? AND user_id = ?", product.StoreID, userID).First(&store).Error; err != nil {
		return errors.New("unauthorized")
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return errors.New("failed to delete product")
	}

	return nil
}

func GetMyProducts(userID uint) ([]model.Product, error) {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, errors.New("store not found")
	}

	var products []model.Product
	if err := config.DB.
		Preload("Store").
		Preload("Category").
		Where("store_id = ?", store.ID).
		Find(&products).Error; err != nil {
		return nil, errors.New("failed to fetch products")
	}

	return products, nil
}
