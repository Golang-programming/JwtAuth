package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	authGroup.POST("/register", RegisterController)
	authGroup.POST("/login", LoginController)
	authGroup.POST("/refresh", RefreshTokenController)

	// Apply authentication middleware to routes
	// router.Use(middleware.AuthMiddleware())
}
