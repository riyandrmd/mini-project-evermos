package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func StoreRoutes(app *fiber.App) {
	group := app.Group("/store", middleware.AuthJWT())
	group.Get("/", handler.GetStore)
	group.Put("/", handler.UpdateStore)
}
