package user

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"title"`
	LastName  string `json:"author"`
}

func MigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&User{})

	if err != nil {
		log.Fatal(err)
	}
}
