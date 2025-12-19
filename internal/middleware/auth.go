package middleware

import (
	"expense-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {

	if c.Request.URL.Path == "/user/login" || c.Request.URL.Path == "/user/signup" {
		c.Next()
		return
	}

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	username, err := utils.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.Set("username", username)

	c.Next()
}
