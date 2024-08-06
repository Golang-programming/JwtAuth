package models

import (
	authModels "blog-app/auth/models"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title         string `gorm:"type:varchar(255); NOT NULL" json:"title"`
	Content       string `gorm:"type:text; NOT NULL" json:"content"`
	ThumbnailUrl  string `gorm:"type:varchar(255); NOT NULL" json:"thumbnail_url"`
	BackgroundUrl string `gorm:"type:varchar(255); NULL" json:"background_url"`

	UserID uint            `gorm:"NOT NULL" json:"user_id"` // Foreign key
	User   authModels.User `gorm:"foreignKey:UserID"`
}
