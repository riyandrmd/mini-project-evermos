package service

import (
	"errors"
	"toko-api/config"
	"toko-api/model"
)

func CreateAlamat(userID uint, input model.Alamat) (*model.Alamat, error) {
	input.UserID = userID
	if err := config.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func GetAlamatByUser(userID uint) ([]model.Alamat, error) {
	var alamat []model.Alamat
	if err := config.DB.Where("user_id = ?", userID).Find(&alamat).Error; err != nil {
		return nil, err
	}
	return alamat, nil
}

func GetAlamatByID(userID uint, id string) (*model.Alamat, error) {
	var alamat model.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}
	if alamat.UserID != userID {
		return nil, errors.New("alamat bukan milik Anda")
	}
	return &alamat, nil
}

func UpdateAlamat(userID uint, id string, input model.Alamat) (*model.Alamat, error) {
	alamat, err := GetAlamatByID(userID, id)
	if err != nil {
		return nil, err
	}
	// update field (kecuali id & user_id)
	alamat.JudulAlamat = input.JudulAlamat
	alamat.NamaPenerima = input.NamaPenerima
	alamat.NoTelp = input.NoTelp
	alamat.DetailAlamat = input.DetailAlamat
	alamat.IDProvinsi = input.IDProvinsi
	alamat.IDKota = input.IDKota

	if err := config.DB.Save(&alamat).Error; err != nil {
		return nil, err
	}
	return alamat, nil
}

func DeleteAlamat(userID uint, id string) error {
	alamat, err := GetAlamatByID(userID, id)
	if err != nil {
		return err
	}
	return config.DB.Delete(&alamat).Error
}
