package blog

import (
	"blog-app/auth/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/api/blogs", GetBlogsController)
	router.POST("/api/blogs", CreateBlogController)
}
