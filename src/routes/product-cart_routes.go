package routes

import (
	"synapsis-rest/src/controllers"
	"synapsis-rest/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductCartRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	productCart := api.Group("/product-cart")
	productCart.Post("/", middleware.Protected(), controllers.CreateProductCart)
	productCart.Get("/", middleware.Protected(), controllers.GetAllProductCart)
}
