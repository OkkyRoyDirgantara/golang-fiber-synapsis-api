package routes

import (
	"synapsis-rest/src/controllers"
	"synapsis-rest/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	routes := api.Group("/transaction")
	routes.Post("/", middleware.Protected(), controllers.CreateTransaction)
}
