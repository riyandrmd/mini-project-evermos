package middleware

import (
	"toko-api/config"
	"toko-api/model"
	"toko-api/utils"

	"github.com/gofiber/fiber/v2"
)

func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		if userID == nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		var user model.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User not found")
		}

		if !user.IsAdmin {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Admin access required")
		}

		return c.Next()
	}
}
