package handler

import (
	"strconv"
	"toko-api/dto"
	"toko-api/service"
	"toko-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateTrx(c *fiber.Ctx) error {
	var req dto.CreateTrxRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	if err := service.CreateTrx(userID, req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Transaction created successfully", nil)
}

func GetMyTransactions(c *fiber.Ctx) error {
	userIDFloat := c.Locals("user_id")
	userID, ok := userIDFloat.(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	trxList, err := service.GetMyTransactions(uint(userID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch transactions")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Transactions fetched", trxList)
}

func GetTrxByID(c *fiber.Ctx) error {
	userIDFloat := c.Locals("user_id")
	userID, ok := userIDFloat.(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	idParam := c.Params("id")
	trxID, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid transaction ID")
	}

	trx, err := service.GetTrxByID(uint(userID), uint(trxID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Transaction fetched", trx)
}
