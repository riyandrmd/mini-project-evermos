package service

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateStore(userID uint, storeName string) error {
	store := model.Store{
		Name:   storeName,
		UserID: userID,
	}
	return config.DB.Create(&store).Error
}

func GetStoreByUserID(userID uint) (*model.Store, error) {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func UpdateStoreByUserID(userID uint, input dto.UpdateStoreRequest) (*model.Store, error) {
	var store model.Store

	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return nil, err
	}

	store.Name = input.Name
	if err := config.DB.Save(&store).Error; err != nil {
		return nil, err
	}

	return &store, nil
}
