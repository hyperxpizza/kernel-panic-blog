package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
)

func main() {

	database.InitDB()

	//Get port from .env file
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8888"
	}

	router := gin.Default()

	router.Run(port)

}
