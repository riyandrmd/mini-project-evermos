package main

import (
	"log"
	"toko-api/config"
	"toko-api/handler"
	"toko-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	config.InitDB()

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Databases connect successfuly")
	// })
	app.Post("/auth/register", handler.Register)
	app.Post("/auth/login", handler.Login)

	category := app.Group("/category")
	category.Get("/", handler.GetCategories)

	category.Use(middleware.AuthJWT(), middleware.RequireAdmin())
	category.Post("/", handler.CreateCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)

	log.Fatal(app.Listen(":8080"))
}
