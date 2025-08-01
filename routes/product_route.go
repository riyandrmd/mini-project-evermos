package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router) {
	product := api.Group("/products", middleware.AuthJWT())

	product.Post("/", handler.CreateProduct)
	product.Get("/", handler.GetMyProducts)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)
}
