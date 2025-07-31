package routes

import (
	"toko-api/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	group := app.Group("/auth")
	group.Post("/register", handler.Register)
	group.Post("/login", handler.Login)
}
