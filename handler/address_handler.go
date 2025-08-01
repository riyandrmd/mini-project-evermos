package handler

import (
	"strconv"
	"time"
	"toko-api/dto"
	"toko-api/model"
	"toko-api/service"
	"toko-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func toAddressResponse(a *model.Address) dto.AddressResponse {
	return dto.AddressResponse{
		ID:        a.ID,
		Label:     a.Label,
		Recipient: a.Recipient,
		Phone:     a.Phone,
		Address:   a.Address,
		UserID:    a.UserID,
		CreatedAt: a.CreatedAt.Format(time.RFC3339),
		UpdatedAt: a.UpdatedAt.Format(time.RFC3339),
	}
}

func CreateAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	var input dto.CreateAddressRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := validator.New().Struct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	address, err := service.CreateAddress(userID, input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Address created", toAddressResponse(address))
}

func GetMyAddresses(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	addresses, err := service.GetAddressesByUser(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var result []dto.AddressResponse
	for _, addr := range addresses {
		result = append(result, toAddressResponse(&addr))
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Addresses fetched", result)
}

func UpdateAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid address ID")
	}

	var input dto.CreateAddressRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := validator.New().Struct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	addr, err := service.UpdateAddress(userID, uint(id), input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Address updated", toAddressResponse(addr))
}

func DeleteAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid address ID")
	}

	if err := service.DeleteAddress(userID, uint(id)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Address deleted", nil)
}
