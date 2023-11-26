package controllers

import (
	"log"
	"synapsis-rest/src/database"
	"synapsis-rest/src/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {
	type NewProduct struct {
		Name  string  `json:"fullname"`
		Price float64 `json:"price"`
		Stock int64   `json:"stock"`
	}

	db := database.DB

	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Println(product)

	if err := db.Create(&product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create product",
			"data":    err,
		})
	}

	newProduct := NewProduct{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Product created",
		"data":    newProduct,
	})
}

func GetAllProduct(ctx *fiber.Ctx) error {
	db := database.DB
	var product []models.Product
	db.Find(&product)
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "All product",
		"data":    product,
	})
}
