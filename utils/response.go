package utils

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, status int, message string, data any) error {
	type Success struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	response := Success{
		Success: true,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	type Error struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	response := Error{
		Success: false,
		Message: message,
	}

	return c.Status(status).JSON(response)
}
