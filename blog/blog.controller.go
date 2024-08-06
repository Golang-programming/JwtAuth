package blog

import (
	"blog-app/blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlogController(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBlog, err := CreateBlogService(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdBlog)
}

func GetBlogsController(c *gin.Context) {
	blogs, err := GetAllBlogsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}
