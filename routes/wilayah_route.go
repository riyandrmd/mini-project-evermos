package routes

import (
	"toko-api/handler"

	"github.com/gofiber/fiber/v2"
)

func WilayahRoutes(app *fiber.App) {
	app.Get("/provinces", handler.GetProvinces)
}
