package service

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input dto.RegisterRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Phone:    input.Phone,
		Address:  input.Address,
		Gender:   input.Gender,
		IsAdmin:  false,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
