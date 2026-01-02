package cmd

import (
	databse "expense-tracker/internal/db"
	"expense-tracker/internal/handler/auth"
	"expense-tracker/internal/handler/expense"
	repository2 "expense-tracker/internal/repository"
	"expense-tracker/internal/service"

	_ "expense-tracker/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						ExpenseTracker
// @description					Manage user expenses
// @termsOfService				http://swagger.io/terms/
// @contact.name 				Dev Support
// @contact.url 				http://www.swagger.io/support
// @securityDefinitions.apikey 	BearerAuth
// @in                         	header
// @name                       	Authorization
// @contact.email 				support@swagger.io
// @license.name 				MIT
// @license.url 				https://opensource.org/licenses/MIT
// @host 						localhost:8088
// @BasePath 					/api/v1
func Run() {
	db := databse.NewDB()

	userRepo := repository2.NewUserRepository(db)
	expenseRepo := repository2.NewExpenseRepository(db)
	userService := service.NewUserService(userRepo, expenseRepo)

	authHandler := auth.NewAuthHandler(userService)
	expenseHandler := expense.NewExpenseHandler(userService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rg := router.Group("/api/v1")

	auth.RegisterRoutes(rg.Group("/auth"), authHandler)
	expense.RegisterRoutes(rg.Group("/expense"), expenseHandler)

	router.Run(":8088")

}
