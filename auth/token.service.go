package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(userID, email string) (string, string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)

	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, _ := accessToken.SignedString(jwtSecret)
	refreshTokenStr, _ := refreshToken.SignedString(jwtSecret)

	// store refresh token in redis

	return accessTokenStr, refreshTokenStr, nil
}

func DecodeToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims, err := DecodeToken(tokenString)

	if err != nil {
		return nil, err
	}
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func RefreshToken(oldAccessToken string) (string, string, error) {

	claims, err := ValidateToken(oldAccessToken)
	fmt.Println("SSS", claims.UserID)

	// check and get token from redis

	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, _ := CreateToken(claims.UserID, claims.Email)

	return accessToken, refreshToken, nil
}

func RevokeToken(tokenString string) error {
	// Implement token revocation logic if needed
	// For example, adding token to a blacklist or removing from a store
	return nil
}
