package utils

import (
	"fmt"
	"hrms-auth-service/internal/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func IsValidToken(tokenString string) bool {
	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(configs.JWTSecret), nil
	})

	if err != nil {
			fmt.Println("JWT parsing error:", err)
			return false
	}

	if !token.Valid {
			return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
			return false
	}

	if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
					return false
			}
	}

	return true
}