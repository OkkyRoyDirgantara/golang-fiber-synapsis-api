package controllers

import (
	"synapsis-rest/src/database"
	"synapsis-rest/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllPost(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.DB.Find(&posts)

	return c.Status(200).JSON(posts)
}
