package database

import (
	"fmt"
	"log"
	"os"
	"warteg-system-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		user, pass, host, port, dbname)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: Failed to connect to database: %v", err)
		return
	}

	fmt.Println("Database connection established")

	// Auto Migration
	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Wallet{}, &models.Transaction{})
	if err != nil {
		log.Printf("ERROR: Failed to migrate database: %v", err)
		return
	}

	DB = db
	SeedAdmin()
}

func SeedAdmin() {
	var user models.User
	result := DB.Where("email = ?", "admin@admin.com").First(&user)
	
	if result.Error != nil {
		// Admin not found, create it
		hashedPassword, _ := models.HashPassword("admin123")
		admin := models.User{
			Username: "admin",
			Email:    "admin@admin.com",
			Password: hashedPassword,
			FullName: "System Administrator",
			Phone:    "0000000000",
		}
		DB.Create(&admin)
		log.Println("Default admin account created: admin / admin123")
	}
}
