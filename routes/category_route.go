package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	group := app.Group("/categories")

	group.Get("/", handler.GetAllCategories)
	group.Post("/", middleware.AuthJWT(), handler.CreateCategory)
	group.Put("/:id", middleware.AuthJWT(), handler.UpdateCategory)
	group.Delete("/:id", middleware.AuthJWT(), handler.DeleteCategory)
}
