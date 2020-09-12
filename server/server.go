package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/handlers"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
)

func main() {

	database.InitDB()

	//Get port from .env file
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8888"
	}

	router := gin.Default()

	//unprotected routes
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	//router.GET("/api/posts", handlers.GetAllPosts)

	//protected routes
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", handlers.GetAllUsers)
	}

	router.Run(port)

}
