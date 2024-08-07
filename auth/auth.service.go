package auth

import (
	"blog-app/auth/models"
	"blog-app/database"

	"golang.org/x/crypto/bcrypt"
)

func registerUserService(input *RegisterInput) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Username: input.Username, Password: string(hashedPassword), Email: input.Email}
	result := database.DB.Create(&user)
	return user, result.Error
}

func authenticateUserService(input *LoginInput) (models.User, error) {
	var user models.User
	result := database.DB.Where("username = ? OR email = ?", input.Identifier, input.Identifier).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, err
	}
	return user, nil
}
