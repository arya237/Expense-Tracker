package routes

import (
	"expense-tracker/controllers"
	"expense-tracker/middleware"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/validator/v10"
)

func SetupRoutes(engine *gin.Engine){

	user := engine.Group("/user")

	user.POST("/signup", controllers.Signup)
	user.POST("/login", controllers.Login)

	expense := engine.Group("/expense")
	expense.Use(middleware.Authorization)

	expense.POST("/addExpense", controllers.AddExpense)
	expense.GET("/updateExpense", controllers.UpdateExpenseStatus)
	expense.GET("/listExpenseByDate", controllers.ListExpenseByDate)
	expense.GET("/listExpenseByDeadLine/:date", controllers.ListExpenseByDeadLine)
	expense.GET("/deleteExpense", controllers.DeleteExpense)
}

func StartServer(engine *gin.Engine, listenAddress string) error{
	err := engine.Run(listenAddress)
	return err 
}