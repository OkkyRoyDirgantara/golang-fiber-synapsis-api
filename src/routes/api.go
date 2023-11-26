package routes

import "github.com/gofiber/fiber/v2"

func ApiRoutes(app *fiber.App) {
	PostRoutes(app)
	RootRoutes(app)
	CustomerRoutes(app)
	AuthRoutes(app)
	ProductRoutes(app)
	ProductCartRoutes(app)
}
