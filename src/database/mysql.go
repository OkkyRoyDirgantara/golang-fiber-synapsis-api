package database

import (
	"fmt"
	"log"
	"os"
	"synapsis-rest/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlConnect() {
	dbUser := "root"
	dbPass := ""
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "synapsis"

	// Buat string koneksi tanpa DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Post{}, &models.Customer{}, &models.Product{}, &models.ProductCart{})

	if err != nil {
		log.Fatal("Connection Refused : \n", err)
		os.Exit(2)
	}

	log.Println("Connection established")
	DB = db
}
