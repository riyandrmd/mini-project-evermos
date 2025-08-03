package routes

import (
	"toko-api/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}
