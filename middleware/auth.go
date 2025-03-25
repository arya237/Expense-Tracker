package middleware

import (
	"expense-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization(c *gin.Context){

	if c.Request.URL.Path == "/user/login" || c.Request.URL.Path == "/user/signup"{
		c.Next()
		return
	}

	tokenString := c.GetHeader("Authorization")

	if tokenString == ""{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return 
	}

	claims := &utils.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return utils.JwtKey, nil
	})

	if err != nil || !token.Valid{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is incorrect"})
		return
	}

	c.Next()
}