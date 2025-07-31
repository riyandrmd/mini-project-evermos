package service

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateCategory(input dto.CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{Name: input.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func UpdateCategory(id string, input dto.UpdateCategoryRequest) (*model.Category, error) {
	var category model.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	category.Name = input.Name
	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func DeleteCategory(id string) error {
	if err := config.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
