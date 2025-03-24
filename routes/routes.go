package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine){

	user := engine.Group("/user")
	user.POST("/signup", controllers.Signup)
	user.POST("/login", controllers.Login)

}

func StartServer(engine *gin.Engine, listenAddress string) error{
	err := engine.Run(listenAddress)
	return err 
}