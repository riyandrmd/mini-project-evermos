package service

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateCategory(req dto.CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{Name: req.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func UpdateCategory(id uint, req dto.UpdateCategoryRequest) (*model.Category, error) {
	var category model.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	category.Name = req.Name
	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func DeleteCategory(id uint) error {
	return config.DB.Delete(&model.Category{}, id).Error
}
