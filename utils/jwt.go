package utils

import (
	"drywave/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "k87qg9ZgjpHPpaA5gH7Upx"

func CreateNewToken(userToken models.UserToken) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userToken)
	return accessToken.SignedString([]byte(SecretKey))
}

func CheckToken(tokenString string) (*models.UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (any, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		fmt.Println("Error parsing/verifying token:", err)
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	dataFromToken, ok := token.Claims.(*models.UserToken)
	if !ok {
		return nil, fmt.Errorf("error extracting claims")
	}
	return dataFromToken, nil
}
