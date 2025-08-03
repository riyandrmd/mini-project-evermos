package main

import (
	"log"
	"toko-api/config"
	"toko-api/routes"

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
	api := app.Group("/api/v1")

	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	routes.AddressRoutes(api)
	routes.StoreRoutes(api)
	routes.CategoryRoutes(api)
	routes.ProductRoutes(api)
	routes.TrxRoutes(api)
	routes.ProvinsiRoutes(api)

	log.Fatal(app.Listen(":3000"))
}
