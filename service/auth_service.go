package service

import (
	"errors"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input dto.RegisterRequest) error {
	var existingUser model.User
	if err := config.DB.Where("email = ? OR notelp = ?", input.Email, input.Notelp).First(&existingUser).Error; err == nil {
		return errors.New("email atau no telp sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal mengenkripsi kata sandi")
	}

	user := model.User{
		Nama:         input.Nama,
		KataSandi:    string(hashedPassword),
		Notelp:       input.Notelp,
		Email:        input.Email,
		TanggalLahir: input.TanggalLahir,
		JenisKelamin: input.JenisKelamin,
		Tentang:      input.Tentang,
		Pekerjaan:    input.Pekerjaan,
		IDProvinsi:   input.IDProvinsi,
		IDKota:       input.IDKota,
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		return err
	}

	toko := model.Toko{
		NamaToko:  "Toko " + input.Nama,
		Deskripsi: "Toko milik " + input.Nama,
		UserID:    user.ID,
	}

	err = config.DB.Create(&toko).Error
	if err != nil {
		config.DB.Unscoped().Delete(&user)
		return errors.New("gagal membuat toko")
	}

	return nil
}
