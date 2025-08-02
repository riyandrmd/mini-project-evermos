package handler

import (
	"toko-api/config"
	"toko-api/dto"
	"toko-api/model"
	"toko-api/service"
	"toko-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var input dto.RegisterRequest

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.New().Struct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var existing model.User
	if err := config.DB.Where("email = ? OR phone = ?", input.Email, input.Phone).First(&existing).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, "Email or phone already exists")
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashed,
		Phone:    input.Phone,
		Gender:   input.Gender,
		Address:  input.Address,
		IsAdmin:  false,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	store := model.Store{
		Name:   user.Name,
		UserID: user.ID,
	}
	if err := config.DB.Create(&store).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create store")
	}

	response := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", response)
}

func Login(c *fiber.Ctx) error {
	var input dto.LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	token, err := service.LoginUser(input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Login Successful", dto.LoginResponse{
		Token: token,
	})
}
