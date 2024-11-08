package database

import (
	"log"

	"github.com/harshtricon/notes-backend-main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/mynotes?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	DB = connection
	if err := connection.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatalf("could not migrate database schema: %v", err)
	}

	log.Println("Connected to MySQL database successfully.")
}
