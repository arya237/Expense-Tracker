package controllers

import (
	"expense-tracker/models"
	"expense-tracker/utils"
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

func Login(c *gin.Context){
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil{
		log.Print("error to parsing user: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	
	hash, _ := utils.HashPassword("1234")
	err = utils.CompareHashedPassword(hash, user.Password);

	if  user.Username != "arya237" || err != nil{
		log.Print("login was unsuccessful: ", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username/password is incorrect"})
		return 
	}

	c.JSON(http.StatusOK, user)
}