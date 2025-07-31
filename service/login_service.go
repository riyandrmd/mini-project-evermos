package service

import (
	"errors"
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
	"toko-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(input dto.LoginRequest) (string, error) {
	var user model.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}
