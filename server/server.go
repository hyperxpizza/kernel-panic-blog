package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/handlers"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//conect to the database
	database.InitDB()

	//create router
	router := gin.Default()
	router.Use(cors.Default())

	//unprotected routes
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	router.GET("/posts", handlers.GetAllPosts)
	router.GET("/posts/:slug", handlers.GetPost)
	router.GET("/posts/:id/comments", handlers.GetComments)
	router.POST("/posts/:id/comments", handlers.AddComment)
	router.GET("/posts/:id/tags", handlers.GetTags)
	//protected routes
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", handlers.Logout) //not working
		protected.GET("/claims", handlers.GetClaims)
		protected.POST("/posts", handlers.CreatePost)
		protected.POST("/tags", handlers.AddTag)
	}

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
