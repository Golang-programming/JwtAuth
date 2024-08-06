package blog

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	blogGroup := router.Group("/api/blogs")
	{
		blogGroup.GET("", GetBlogsController)
		blogGroup.POST("", CreateBlogController)
	}
}
