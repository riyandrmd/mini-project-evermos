package handler

import (
	"toko-api/dto"
	"toko-api/service"
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetStore(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	store, err := service.GetStoreByUserID(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Store not found")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Store fetched", dto.StoreResponse{
		ID:        store.ID,
		Name:      store.Name,
		UserID:    store.UserID,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	})
}

func UpdateStore(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid user ID")
	}
	userID := uint(userIDFloat)

	var input dto.UpdateStoreRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	updatedStore, err := service.UpdateStoreByUserID(userID, input)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Store updated", dto.StoreResponse{
		ID:        updatedStore.ID,
		Name:      updatedStore.Name,
		UserID:    updatedStore.UserID,
		CreatedAt: updatedStore.CreatedAt,
		UpdatedAt: updatedStore.UpdatedAt,
	})
}
