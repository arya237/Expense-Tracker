package controllers

import (
	"expense-tracker/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context){
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil{
		log.Print("can't parse user: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"You registered successfuly!"})
}