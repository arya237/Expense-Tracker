package controllers

import (
	"expense-tracker/database"
	"expense-tracker/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context){

	var expense models.Expense
	username := c.Query("username")

	err := c.ShouldBindBodyWithJSON(&expense)

	if err != nil{
		log.Print("can't bind expense: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	expense.UserID = username

	err = database.AddExpenseToDatabase(expense)
	
	if err != nil{
		log.Print("can't add expense to database: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"Expense added successfuly", "expense": expense})
}

func ListExpense(c *gin.Context){
	
	filter := c.Query("filter")
	username := c.Query("username")
	
	list, err := database.ListExpense(filter, username)

	if err != nil{
		log.Print("can't get list of expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, list)
}