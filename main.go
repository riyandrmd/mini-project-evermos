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

	routes.AuthRoutes(app)
	routes.CategoryRoutes(app)
	routes.StoreRoutes(app)
	routes.ProductRoutes(app)
	routes.AddressRoutes(app)
	routes.TrxRoutes(app)
	routes.WilayahRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
