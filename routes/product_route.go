package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	products := app.Group("/products")

	products.Post("/", middleware.AuthJWT(), handler.CreateProduct)
	products.Get("/", handler.GetAllProducts)
	products.Get("/:id", handler.GetProductByID)
	products.Put("/:id", middleware.AuthJWT(), handler.UpdateProduct)
	products.Delete("/:id", middleware.AuthJWT(), handler.DeleteProduct)
}
