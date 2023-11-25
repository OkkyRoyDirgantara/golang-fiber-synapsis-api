package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CustomerRoutes(app *fiber.App) {
	api := app.Group("api/v1", logger.New())

	customer := api.Group("/customer", logger.New())
	customer.Get("/:id", controllers.GetCustomer)
	customer.Post("/", controllers.CreateCustomer)
}
