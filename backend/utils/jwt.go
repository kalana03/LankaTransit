package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

jwtKey := []byte("MySecretKey")

type Claims struct{
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken (userId string)(string, error){
	exp_time := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}