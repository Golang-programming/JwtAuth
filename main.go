package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/test/gingonic/database"
	"github.com/test/gingonic/user"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := gin.Default()
	app.Use(gin.Logger())

	db := database.ConnectToDatabase()

	user.MigrateUser(db)

	repo := database.Repository{
		DB: db,
	}

	app.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Access granted for api-1"})
	})

	app.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Access granted for api-2"})
	})

	app.Run(":", port)

}
