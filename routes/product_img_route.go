package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductImageRoutes(app *fiber.App) {
	product := app.Group("/products", middleware.AuthJWT())
	product.Post("/:id/images", handler.UploadProductImage)
	product.Get("/:id/images", handler.GetProductImages)
	product.Delete("/images/:imageId", handler.DeleteProductImage)

}
