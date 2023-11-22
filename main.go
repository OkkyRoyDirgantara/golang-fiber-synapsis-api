package main

import (
	"synapsis-rest/src/database"
	"synapsis-rest/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	database.MysqlConnect()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.PostRoutes(app)

	app.Listen(":3000")
}
