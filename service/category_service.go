package service

import (
	"toko-api/config"
	"toko-api/model"
)

func GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryByID(id string) (*model.Category, error) {
	var category model.Category
	err := config.DB.First(&category, id).Error
	return &category, err
}

func CreateCategory(input model.Category) (*model.Category, error) {
	err := config.DB.Create(&input).Error
	return &input, err
}

func UpdateCategory(id string, input model.Category) (*model.Category, error) {
	var category model.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	category.NamaCategory = input.NamaCategory
	err := config.DB.Save(&category).Error
	return &category, err
}

func DeleteCategory(id string) error {
	return config.DB.Delete(&model.Category{}, id).Error
}
