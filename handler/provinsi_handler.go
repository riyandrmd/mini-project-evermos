package handler

import (
	"toko-api/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllProvinsi(c *fiber.Ctx) error {
	data, err := service.GetAllProvinsi()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": true, "data": data})
}
