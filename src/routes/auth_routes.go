package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	routes := api.Group("/auth")
	routes.Post("/login", controllers.Login)
}
