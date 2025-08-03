package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func StoreRoutes(app *fiber.App) {
	toko := app.Group("/toko", middleware.AuthJWT())

	toko.Get("/", handler.GetMyToko)
	toko.Put("/", handler.UpdateMyToko)
}
