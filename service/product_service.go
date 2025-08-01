package service

import (
	"errors"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateProduct(userID uint, input dto.CreateProductRequest) (*model.Product, error) {
	var store model.Store

	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, errors.New("you do not own a store")
	}

	var category model.Category
	if err := config.DB.First(&category, input.CategoryID).Error; err != nil {
		return nil, errors.New("category not found")
	}

	product := model.Product{
		Name:       input.Name,
		Price:      input.Price,
		Stock:      input.Stock,
		StoreID:    store.ID,
		CategoryID: input.CategoryID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProductByID(userID, productID uint, input dto.CreateProductRequest) (*model.Product, error) {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, errors.New("unauthorized: store not found")
	}

	var product model.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}

	if product.StoreID != store.ID {
		return nil, errors.New("unauthorized: not your product")
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func DeleteProductByID(userID, productID uint) error {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return errors.New("unauthorized: store not found")
	}

	var product model.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		return errors.New("product not found")
	}

	if product.StoreID != store.ID {
		return errors.New("unauthorized: not your product")
	}

	return config.DB.Delete(&product).Error
}

func GetProductsByUser(userID uint) ([]model.Product, error) {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, errors.New("store not found")
	}

	var products []model.Product
	if err := config.DB.Where("store_id = ?", store.ID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
