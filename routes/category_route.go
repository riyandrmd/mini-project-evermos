package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	c := app.Group("/categories")

	c.Get("/", handler.GetAllCategories)
	c.Get("/:id", handler.GetCategoryByID)

	// admin only
	c.Post("/", middleware.AuthJWT(), middleware.AdminOnly(), handler.CreateCategory)
	c.Put("/:id", middleware.AuthJWT(), middleware.AdminOnly(), handler.UpdateCategory)
	c.Delete("/:id", middleware.AuthJWT(), middleware.AdminOnly(), handler.DeleteCategory)
}
