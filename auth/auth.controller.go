package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerController(c *gin.Context) {
	var registerInput = &RegisterInput{}
	if err := c.BindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("registerInput", registerInput)
	user, err := registerUserService(registerInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func loginController(c *gin.Context) {
	var loginInput = &LoginInput{}
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authenticateUserService(loginInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, user)
}
