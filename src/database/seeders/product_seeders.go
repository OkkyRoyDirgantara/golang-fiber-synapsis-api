package seeders

import (
	"synapsis-rest/src/models"

	"gorm.io/gorm"
)

func ProductSeeder(db *gorm.DB) {
	products := []models.Product{
		{Name: "Semangka", Price: 10, Stock: 10},
		{Name: "Jambu", Price: 10, Stock: 10},
	}

	for _, product := range products {
		db.Create(&product)
	}
}
