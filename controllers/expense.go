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

func ListExpenseByDate(c *gin.Context){
	
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

func ListExpenseByDeadLine(c *gin.Context){
	
	date := c.Param("date")
	username := c.Query("username")

	list, err := database.ListExpenseWithDeadLine(date, username)

	if err != nil{
		log.Print("can't get list of expenses: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please try again!"})
		return 
	}

	c.JSON(http.StatusOK, list)
}

func UpdateExpenseStatus(c *gin.Context){

	expenseID := c.Query("id")
	status := c.Query("status")

	err := database.UpdateStatus(expenseID, status)

	if err != nil{
		log.Print("can't update expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your expense updated successfuly"})
}

func DeleteExpense(c *gin.Context){
	expenseID := c.Query("id")

	err := database.DeleteExpenseFromDatabase(expenseID)

	if err != nil{
		log.Print("can't delete this expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please try again"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your expense deleted successfuly"})
}