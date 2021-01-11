package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

func GetAuthToken(id int, isAdmin bool, username string) (string, error) {
	if tokenSecret == nil {
		log.Println("TokenSecret is nil")
		tokenSecret = []byte("tokenSecrettest")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["username"] = username
	claims["is_admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)

	return authToken, err
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return tokenSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func IsTokenValid(tokenString string) (bool, float64, string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Token signing method is not valid: %v", token.Header["alg"])
		}

		return tokenSecret, nil
	})

	if err != nil {
		fmt.Printf("Err %v \n", err)
		return false, 0, "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["user_id"]
		username := claims["username"]
		isAdmin := claims["is_admin"]
		return true, id.(float64), username.(string), isAdmin.(bool)
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		fmt.Println(err)
		return false, 0, "", false
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}

		tokenString := cookie.Value
		isValid, id, username, isAdmin := IsTokenValid(tokenString)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", id)
			c.Set("username", username)
			c.Set("isAdmin", isAdmin)
			c.Next()
		}
	}
}

func ExtractToken(c *gin.Context) (*string, error) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	return &tokenString, nil
}
