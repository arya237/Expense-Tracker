package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret_key")

type Claims struct{
	jwt.StandardClaims
} 

func CreateJwtClaims() *Claims{
	expirationDate := time.Now().Add(12 * time.Hour)

	return &Claims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationDate.Unix()},
	}
}

func CreateToken(claims *Claims) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil{
		return "", err
	}

	return tokenString, nil
}