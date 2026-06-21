package main

import (
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/handlers"
)

func main() {
	database.ConnectDB()
	router:=gin.Default()
	router.Static("/uploads", "./uploads")
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.CreateProduct)
	router.Run(":8080")
}