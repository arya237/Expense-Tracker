package controllers

import (
	"expense-tracker/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context){

	var expense models.Expense
	err := c.ShouldBindBodyWithJSON(&expense)

	if err != nil{
		log.Print("can't bind expense: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message":"Expense added successfuly"})
}	