package service

import (
	"errors"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
)

func CreateAddress(userID uint, input dto.CreateAddressRequest) (*model.Address, error) {
	addr := model.Address{
		UserID:    userID,
		Label:     input.Label,
		Recipient: input.Recipient,
		Phone:     input.Phone,
		Address:   input.Address,
	}
	err := config.DB.Create(&addr).Error
	return &addr, err
}

func GetAddressesByUser(userID uint) ([]model.Address, error) {
	var addrs []model.Address
	err := config.DB.Where("user_id = ?", userID).Find(&addrs).Error
	return addrs, err
}

func UpdateAddress(userID, addrID uint, input dto.CreateAddressRequest) (*model.Address, error) {
	var updated model.Address

	result := config.DB.Model(&updated).
		Where("id = ? AND user_id = ?", addrID, userID).
		Updates(model.Address{
			Label:     input.Label,
			Recipient: input.Recipient,
			Phone:     input.Phone,
			Address:   input.Address,
		})

	if result.RowsAffected == 0 {
		return nil, errors.New("address not found or not owned by user")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.First(&updated, addrID)
	return &updated, nil
}

func DeleteAddress(userID, addrID uint) error {
	result := config.DB.Where("id = ? AND user_id = ?", addrID, userID).Delete(&model.Address{})
	if result.RowsAffected == 0 {
		return errors.New("address not found or not owned by user")
	}
	return result.Error
}
