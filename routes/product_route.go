package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	group := app.Group("/products", middleware.AuthJWT())
	group.Post("/", handler.CreateProduct)
	group.Get("/", handler.GetAllProducts)
	group.Get("/my-products", handler.GetMyProducts)
	group.Get("/:id", handler.GetProductByID)
	group.Put("/:id", handler.UpdateProduct)
	group.Delete("/:id", handler.DeleteProduct)

}
