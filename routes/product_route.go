package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router) {
	product := api.Group("/products", middleware.AuthJWT())

	product.Post("/", handler.CreateProduct)
	product.Get("/", handler.GetAllProducts)
	product.Get("/my-products", handler.GetMyProducts)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)
	product.Get("/:id/images", handler.UploadProductImage)
	product.Post("/:id/images", handler.UploadProductImage)
	product.Delete("/images/:imageId", handler.DeleteProductImage)

}
