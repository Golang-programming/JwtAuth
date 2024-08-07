package main

import (
	"blog-app/auth"
	"blog-app/blog"
	"blog-app/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	router := gin.New()

	if port == "" {
		port = "8000"
	}

	router.Use(gin.Logger())
	auth.RegisterRoutes(router)
	blog.RegisterRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	err = database.InitializeDB()
	auth.InitRedisClient()
	if err != nil {
		panic("failed to connect database")
	}

	router.Run(":" + port)
}
