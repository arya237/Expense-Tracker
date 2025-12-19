package cmd

import (
	databse "expense-tracker/internal/db"
	"expense-tracker/internal/handler/auth"
	"expense-tracker/internal/handler/expense"
	repository2 "expense-tracker/internal/repository"
	"expense-tracker/internal/service"

	"github.com/gin-gonic/gin"
)

//type Container struct {
//
//}

func Run() {
	db := databse.NewDB()

	userRepo := repository2.NewUserRepository(db)
	expenseRepo := repository2.NewExpenseRepository(db)
	userService := service.NewUserService(userRepo, expenseRepo)

	authHandler := auth.NewAuthHandler(userService)
	expenseHandler := expense.NewExpenseHandler(userService)

	router := gin.Default()

	auth.RegisterRoutes(router.Group("/auth"), authHandler)
	expense.RegisterRoutes(router.Group("/expense"), expenseHandler)

	router.Run(":8088")

}
