package controllers

import (
	"log"
	"strings"
	"synapsis-rest/src/database"
	"synapsis-rest/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateTransaction(ctx *fiber.Ctx) error {
	db := database.DB
	tokenString := ctx.Get("Authorization")
	parts := strings.Fields(tokenString)
	tokenString = parts[1]

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Printf("Error %s", err)
	}

	data := token.Claims.(jwt.MapClaims)
	customerId := data["customer_id"].(float64)

	var productCart []models.ProductCart

	// Fetch product cart data for the given customer
	if err := db.Where("customer_id = ?", customerId).Find(&productCart).Error; err != nil {
		log.Printf("Error fetching product cart: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}

	// Check if the cart is empty
	if len(productCart) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cart is empty. Transaction failed.",
		})
	}

	// If the cart is not empty, proceed with creating the transaction
	if err := db.Delete(&productCart).Error; err != nil {
		log.Printf("Error deleting product cart: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Transaction created",
	})
}
