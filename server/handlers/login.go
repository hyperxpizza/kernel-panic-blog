package handlers

import "github.com/gin-gonic/gin"

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

}
