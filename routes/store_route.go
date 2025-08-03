package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func StoreRoutes(router fiber.Router) {
	toko := router.Group("/toko", middleware.AuthJWT())

	toko.Get("/", handler.GetMyToko)
	toko.Put("/", handler.UpdateMyToko)
}
