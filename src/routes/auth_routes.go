package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("api/v1", logger.New())

	customer := api.Group("/auth", logger.New())
	customer.Post("/login", controllers.Login)
}
