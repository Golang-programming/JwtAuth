package auth

import (
	"blog-app/auth/models"
	"blog-app/database"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(username, password string) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Username: username, Password: string(hashedPassword)}
	result := database.DB.Create(&user)
	return user, result.Error
}

func AuthenticateUserService(username, password string) (models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}
	return user, nil
}
