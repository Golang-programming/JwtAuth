package auth

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

var (
	jwtSecret   = []byte("your_secret_key")
	redisClient *redis.Client
	ctx         = context.Background()
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func CreateToken(userID, email string) (string, string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)

	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenStr, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Store the refresh token in Redis
	err = redisClient.Set(ctx, "refresh_token:"+userID, refreshTokenStr, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	claims, err := decodeToken(tokenStr)
	if err != nil {
		return nil, err
	}
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func RefreshToken(refreshTokenStr string) (string, string, error) {
	claims, err := decodeToken(refreshTokenStr)
	if err != nil {
		return "", "", err
	}

	storedRefreshToken, err := redisClient.Get(ctx, "refresh_token:"+claims.UserID).Result()
	if err != nil || storedRefreshToken != refreshTokenStr {
		return "", "", errors.New("refresh token is invalid or does not match")
	}

	accessToken, newRefreshToken, err := CreateToken(claims.UserID, claims.Email)
	if err != nil {
		return "", "", err
	}

	err = redisClient.Set(ctx, "refresh_token:"+claims.UserID, newRefreshToken, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func RevokeToken(userID string) error {
	err := redisClient.Del(ctx, "refresh_token:"+userID).Err()
	if err != nil {
		return err
	}

	return nil
}

func decodeToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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
