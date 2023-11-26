package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	customer := api.Group("/customer")
	customer.Get("/:id", controllers.GetCustomer)
	customer.Post("/", controllers.CreateCustomer)
}
