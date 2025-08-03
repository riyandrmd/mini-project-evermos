package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func TrxRoutes(router fiber.Router) {
	trx := router.Group("/transactions", middleware.AuthJWT())
	trx.Post("/", handler.CreateTransaction)
	trx.Get("/", handler.GetMyTransactions)
	trx.Get("/:id", handler.GetTransactionByID)
}
