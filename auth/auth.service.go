package auth

import (
	"blog-app/auth/models"
	"blog-app/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(input *RegisterInput) (models.User, string, string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Username: input.Username, Password: string(hashedPassword), Email: input.Email}
	result := database.DB.Create(&user)

	accessToken, refreshToken, err := CreateToken(string(user.ID), user.Email)
	if err != nil {
		return user, "", "", errors.New("token creation failed: " + err.Error())
	}

	return user, accessToken, refreshToken, result.Error
}

func AuthenticateUserService(input *LoginInput) (models.User, string, string, error) {
	var user models.User
	result := database.DB.Where("username = ? OR email = ?", input.Identifier, input.Identifier).First(&user)
	if result.Error != nil {
		return user, "", "", result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, "", "", err
	}

	accessToken, refreshToken, err := CreateToken(string(user.ID), user.Email)
	if err != nil {
		return user, "", "", errors.New("token creation failed: " + err.Error())
	}

	return user, accessToken, refreshToken, nil
}

func RefreshTokenService(refreshTokenStr string) (string, string, error) {
	return RefreshToken(refreshTokenStr)
}
