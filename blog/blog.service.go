package blog

import (
	"blog-app/blog/models"
	"blog-app/database"
)

func CreateBlogService(blog models.Blog) (models.Blog, error) {
	result := database.DB.Create(&blog)
	return blog, result.Error
}

func GetAllBlogsService() ([]models.Blog, error) {
	var blogs []models.Blog
	result := database.DB.Find(&blogs)
	return blogs, result.Error
}
