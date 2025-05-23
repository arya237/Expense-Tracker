package controllers

import (
	"expense-tracker/database"
	"expense-tracker/models"
	"expense-tracker/utils"
	"log"
	"net/http"
	"strings"

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

	pass, err := utils.HashPassword(user.Password)

	if err != nil{
		log.Print("can't hashing password: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please try again"})
		return
	}

	user.Password = pass

	err = database.AddUserToDatabase(user)

	if err != nil{
		if strings.Contains(err.Error(),"duplicate key"){
			c.JSON(http.StatusConflict, gin.H{"error":"this username already exists"})
			return
		} else{

			log.Print("can't insert user in database: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message":"You registered successfuly!"})
}

func Login(c *gin.Context){
	var income models.User
	err := c.ShouldBindBodyWithJSON(&income)

	if err != nil{
		log.Print("error to parsing user: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	
	user, err := database.GetUserFromDatabase(income)

	if err != nil{
		log.Print("login was unsuccessful: ", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims := utils.CreateJwtClaims()
	token, err  := utils.CreateToken(claims)
	
	if err != nil{
		log.Print("can't create jwt claims: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return 
	} 

	c.JSON(http.StatusOK, gin.H{"user": user, "token":token})
}