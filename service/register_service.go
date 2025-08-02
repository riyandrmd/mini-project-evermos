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
		Gender:   input.Gender,
		About:    input.About,
		Job:      input.Job,
		Province: input.Province,
		City:     input.City,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
