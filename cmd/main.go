package main

import (
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/handlers"
)

func main() {
	database.ConnectDB()
	router:=gin.Default()
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.GetProducts)
	router.Run(":8080")
}