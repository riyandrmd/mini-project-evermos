package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(router fiber.Router) {
	user := router.Group("/user", middleware.AuthJWT())

	user.Post("/alamat", handler.CreateAlamat)
	user.Get("/alamat", handler.GetAlamatSaya)
	user.Get("/alamat/:id", handler.GetAlamatByID)
	user.Put("/alamat/:id", handler.UpdateAlamat)
	user.Delete("/alamat/:id", handler.DeleteAlamat)
}
