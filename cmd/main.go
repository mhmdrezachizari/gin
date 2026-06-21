package main

import (
	"backend/database"
	"backend/handlers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()

	router.Static("/uploads", "./uploads")

	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	router.GET("/products", handlers.GetProducts)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/products", handlers.CreateProduct)
	}

	router.Run(":8080")
}