package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = map[string]string{
	"Host":     os.Getenv("DB_HOST"),
	"Port":     os.Getenv("DB_PORT"),
	"Password": os.Getenv("DB_PASS"),
	"User":     os.Getenv("DB_USER"),
	"SSLMode":  os.Getenv("DB_SSLMODE"),
	"DBName":   os.Getenv("DB_NAME"),
}

type Repository struct {
	DB *gorm.DB
}

func ConnectToDatabase() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config["Host"], config["Port"], config["User"], config["Password"], config["DBName"], config["SSLMode"],
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
