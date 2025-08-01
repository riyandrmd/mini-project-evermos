package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func TrxRoutes(app fiber.Router) {
	trx := app.Group("/transactions", middleware.AuthJWT())
	trx.Post("/", handler.CreateTrx)
	trx.Get("/", handler.GetMyTransactions)
	trx.Get("/:id", handler.GetTrxByID)
}
