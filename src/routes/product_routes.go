package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	product := api.Group("/product")
	product.Post("/", controllers.CreateProduct)
	product.Get("/", controllers.GetAllProduct)
}
