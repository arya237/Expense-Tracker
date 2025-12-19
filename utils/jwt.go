package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret_key")

type Claims struct {
	jwt.StandardClaims
	Username string
}

func CreateJwtClaims(username string) *Claims {
	expirationDate := time.Now().Add(12 * time.Hour)

	return &Claims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationDate.Unix()},
		Username:       username,
	}
}

func CreateToken(claims *Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims.Username, nil
}
