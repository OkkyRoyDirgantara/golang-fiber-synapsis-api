package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RootRoutes(rt *fiber.App) {
	route := rt

	route.Get("/", controllers.GetRoot)
}
