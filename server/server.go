package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/handlers"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
)

func main() {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")

	if user == "" {
		user = "kernelpanicuser"
	}

	if password == "" {
		password = "testkernel"
	}

	if dbname == "" {
		dbname = "kernelpanicblog"
	}

	database.InitDB(user, password, dbname)

	//Get port from .env file
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8888"
	}

	router := gin.Default()

	//use default cors settings
	router.Use(cors.Default())

	//unprotected routes
	router.POST("/login", handlers.Login)       //works
	router.POST("/register", handlers.Register) //works
	router.GET("/posts", handlers.GetAllPosts)  //works
	router.GET("/post/:slug", handlers.GetPostBySlug)
	router.POST("/comment/add", handlers.AddComment)

	//protected routes
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		// test extracting claims
		protected.GET("/claims", handlers.GetClaims) //works

		// users
		protected.GET("/users", handlers.GetAllUsers)

		//posts
		protected.POST("/create/post", handlers.CreatePost)
		protected.POST("/update/post/:id", handlers.UpdatePost)
		protected.DELETE("/delete/post/:id", handlers.DeletePost)
	}

	router.Run(port)

}
