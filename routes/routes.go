package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine){

	user := engine.Group("/user")
	user.POST("/signup", controllers.Signup)
	user.POST("/login", controllers.Login)

	expense := engine.Group("/expense")
	expense.POST("/addExpense", controllers.AddExpense)

}

func StartServer(engine *gin.Engine, listenAddress string) error{
	err := engine.Run(listenAddress)
	return err 
}