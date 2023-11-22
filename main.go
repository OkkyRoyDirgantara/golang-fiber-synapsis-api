package main

import (
	"log"
	"synapsis-rest/src/database"
	"synapsis-rest/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	app := Setup()
	log.Fatal(app.Listen(":3000"))
}

func Setup() *fiber.App {
	engine := mustache.New("./src/views", ".mustache")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(logger.New())

	database.MysqlConnect()

	routes.RootRoutes(app)
	routes.PostRoutes(app)
	return app
}
