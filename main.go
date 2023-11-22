package main

import (
	"synapsis-rest/src/database"
	"synapsis-rest/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	database.MysqlConnect()
	engine := mustache.New("./src/views", ".mustache")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	routes.RootRoutes(app)
	routes.PostRoutes(app)

	app.Listen(":3000")
}
