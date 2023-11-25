package controllers

import (
	"synapsis-rest/src/database"
	"synapsis-rest/src/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var customer models.Customer
	db.Find(&customer, id)
	if customer.Email == "" {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "error",
			"message": "No user found with ID",
			"data":    nil})
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "User found",
		"data":    customer})
}

func CreateCustomer(c *fiber.Ctx) error {
	type NewCustomer struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}

	db := database.DB
	customer := new(models.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(customer.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	customer.Password = hash
	if err := db.Create(&customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newCustomer := NewCustomer{
		Email:    customer.Email,
		Fullname: customer.Fullname,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newCustomer})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
