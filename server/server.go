package main

import (
	"log"
	"os"

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

	router.Use(CORSMiddleware())

	//unprotected routes
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	router.GET("/posts", handlers.GetAllPosts)

	router.GET("/posts/:slug", handlers.GetPost)
	router.GET("/post/:id/comments", handlers.GetComments)
	router.POST("/post/:id/comments", handlers.AddComment)
	//router.GET("/post/:id/tags", handlers.GetTags)

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
