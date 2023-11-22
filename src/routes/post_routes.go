package routes

import (
	"synapsis-rest/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(a *fiber.App) {
	route := a.Group("api/v1")

	route.Get("post", controllers.GetAll)
}
