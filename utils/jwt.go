package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

type Claims struct{
	ID int
	jwt.StandardClaims
} 

func CreateJwtClaims(id int) *Claims{
	expirationDate := time.Now().Add(12 * time.Hour)

	return &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationDate.Unix()},
	}
}

func CreateToken(claims *Claims) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil{
		return "", err
	}

	return tokenString, nil
}