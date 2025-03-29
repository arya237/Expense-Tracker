package main

import (
	"expense-tracker/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	
	router := gin.Default()
	routes.SetupRoutes(router)
	log.Print(routes.StartServer(router, ":8080"))
}