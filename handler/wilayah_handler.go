package handler

import (
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetProvinces(c *fiber.Ctx) error {
	provinces, err := utils.FetchProvinces()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch provinces")
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Provinces fetched", provinces)
}
