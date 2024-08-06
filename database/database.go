package database

import (
	userModel "blog-app/auth/models"
	blogModel "blog-app/blog/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres" // or "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() error {
	var config = map[string]string{
		"Host":     os.Getenv("DB_HOST"),
		"Port":     os.Getenv("DB_PORT"),
		"Password": os.Getenv("DB_PASS"),
		"User":     os.Getenv("DB_USER"),
		"SSLMode":  os.Getenv("DB_SSLMODE"),
		"DBName":   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config["Host"], config["Port"], config["User"], config["Password"], config["DBName"], config["SSLMode"],
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&blogModel.Blog{})

	return nil
}
