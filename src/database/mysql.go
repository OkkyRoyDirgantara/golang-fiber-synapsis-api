package database

import (
	"fmt"
	"log"
	"os"
	"synapsis-rest/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlConnect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default values.")
	}

	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASS", "")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "synapsis")

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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
