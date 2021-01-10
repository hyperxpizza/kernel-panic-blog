package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goware/emailx"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
	Email     string `json:"email"`
}

func Register(c *gin.Context) {
	var registerData RegisterData

	//Parse json into register struct
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	//Validate data
	err := emailx.Validate(registerData.Email)
	if err != nil {
		if err == emailx.ErrInvalidFormat {
			c.JSON(http.StatusConflict, gin.H{
				"status": "Wrong email format",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error while validating email",
		})
		return
	}

	//TODO: add username validation

	//Check if username is already taken
	usernameExists := database.CheckIfUsernameExists(registerData.Username)
	if usernameExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This username is already taken",
		})

		return
	}

	//Check if email is already taken
	emailExists := database.CheckIfEmailTaken(registerData.Email)
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This email is already taken",
		})

		return
	}

	//Check if passwords match
	if registerData.Password1 != registerData.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords do not match",
		})
		return
	}

	//Create password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password1), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while hashing the password",
		})

		return
	}

	registerData.Password1 = string(hashedPassword)

	// Insert into the database
	err = database.InsertUser(registerData.Username, registerData.Password1, registerData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error while inserting user into the database",
		})

		return
	}

	c.JSON(http.StatusOK, &registerData)
}
