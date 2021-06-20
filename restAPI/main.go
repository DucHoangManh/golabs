package main

import (
	"github.com/DucHoangManh/golabs/restAPI/database"
	"github.com/DucHoangManh/golabs/restAPI/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	routes.Setup(app)
	app.Listen(":8000")
}
