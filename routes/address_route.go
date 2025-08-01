package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(app *fiber.App) {
	group := app.Group("/address", middleware.AuthJWT())
	group.Post("/", handler.CreateAddress)
	group.Get("/", handler.GetMyAddresses)
	group.Put("/:id", handler.UpdateAddress)
	group.Delete("/:id", handler.DeleteAddress)
}
