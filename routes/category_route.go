package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(router fiber.Router) {
	category := router.Group("/categories")

	category.Get("/", handler.GetAllCategories)
	category.Get("/:id", handler.GetCategoryByID)

	// admin only
	category.Post("/", middleware.AuthJWT(), middleware.AdminOnly(), handler.CreateCategory)
	category.Put("/:id", middleware.AuthJWT(), middleware.AdminOnly(), handler.UpdateCategory)
	category.Delete("/:id", middleware.AuthJWT(), middleware.AdminOnly(), handler.DeleteCategory)
}
