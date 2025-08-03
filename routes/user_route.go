package routes

import (
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Get("/me", middleware.AuthJWT(), handler.GetMyProfile)
	user.Put("/me", middleware.AuthJWT(), handler.UpdateMyProfile)

}
