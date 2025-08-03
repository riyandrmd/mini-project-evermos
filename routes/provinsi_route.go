package routes

import (
	"toko-api/handler"

	"github.com/gofiber/fiber/v2"
)

func ProvinsiRoutes(router fiber.Router) {
	router.Get("/provinces", handler.GetAllProvinsi)
}
