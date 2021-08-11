package main

import (
	"example.com/user/hello/databases"
	"example.com/user/hello/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	databases.Connect()
	app := fiber.New()
	app.Use(cors.New())
	routes.Setup(app)
	app.Listen(":8000")
}
