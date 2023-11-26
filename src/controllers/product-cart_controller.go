package controllers

import (
	"fmt"
	"strings"
	"synapsis-rest/src/database"
	"synapsis-rest/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProductCart(ctx *fiber.Ctx) error {
	type NewProductCart struct {
		Quantity   int `json:"quantity"`
		CustomerId uint
		ProductId  uint `json:"productId"`
	}

	tokenString := ctx.Get("Authorization")
	parts := strings.Fields(tokenString)
	tokenString = parts[1]

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	data := token.Claims.(jwt.MapClaims)
	customerId := data["customer_id"].(float64)

	db := database.DB
	productCart := new(models.ProductCart)
	if err := ctx.BodyParser(productCart); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(
		&models.ProductCart{
			Quantity:   productCart.Quantity,
			CustomerId: uint(customerId),
			ProductId:  productCart.ProductId,
		},
	).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create product cart",
			"data":    err,
		})
	}

	newProductCart := NewProductCart{
		Quantity:   productCart.Quantity,
		CustomerId: uint(customerId),
		ProductId:  productCart.ProductId,
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Product cart created",
		"data":    newProductCart,
	})
}

func GetAllProductCart(ctx *fiber.Ctx) error {
	db := database.DB
	var productCart []models.ProductCart
	db.Preload("Product").Preload("Customer").Find(&productCart)
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "All product cart",
		"data":    productCart,
	})
}
