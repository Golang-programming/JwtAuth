package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerController(c *gin.Context) {
	// var registerInput = RegisterInput
	var input struct {
		ID    int    `form:"id" binding:"required"`
		Name  string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required,email"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := registerUserService(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func loginController(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := authenticateUserService(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, user)
}
