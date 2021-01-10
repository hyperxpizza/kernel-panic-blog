package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	//Parse JSON into login struct
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	usernameExists := database.CheckIfUsernameExists(loginData.Username)
	if !usernameExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
		})
		return
	}

	//compare passwords
	passwordToCheck := database.GetUsersPassword(loginData.Username)
	if passwordToCheck == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while checking password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordToCheck), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	isAdmin, id := database.GetAdminAndID(loginData.Username)

	//generate token
	token, err := middleware.GetAuthToken(id, isAdmin, loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while logging in",
		})
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged in succesfully",
		"token":   token,
	})

	return

}

func GetClaims(c *gin.Context) {
	token, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims, ok := middleware.ExtractClaims(*token)
	if ok {
		c.JSON(http.StatusOK, &claims)
	}

}
