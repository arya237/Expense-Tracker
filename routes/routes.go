package routes

import (
	"expense-tracker/controllers"
	"expense-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine){

	user := engine.Group("/user")

	user.POST("/signup", controllers.Signup)
	user.POST("/login", controllers.Login)

	expense := engine.Group("/expense")
	expense.Use(middleware.Authorization)

	expense.POST("/addExpense", controllers.AddExpense)
	expense.GET("/updateExpense", controllers.UpdateExpenseStatus)
	expense.GET("/listExpense", controllers.ListExpense)
	expense.GET("/deleteExpense", controllers.DeleteExpense)
}

func StartServer(engine *gin.Engine, listenAddress string) error{
	err := engine.Run(listenAddress)
	return err 
}