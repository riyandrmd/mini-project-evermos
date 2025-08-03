package service

import (
	"errors"
	"toko-api/config"
	"toko-api/model"
)

func GetTokoByUser(userID uint) (*model.Toko, error) {
	var toko model.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return nil, errors.New("toko tidak ditemukan")
	}
	return &toko, nil
}

func UpdateToko(userID uint, input model.Toko) (*model.Toko, error) {
	toko, err := GetTokoByUser(userID)
	if err != nil {
		return nil, err
	}

	toko.NamaToko = input.NamaToko
	toko.Deskripsi = input.Deskripsi
	toko.Foto = input.Foto

	if err := config.DB.Save(&toko).Error; err != nil {
		return nil, err
	}
	return toko, nil
}
