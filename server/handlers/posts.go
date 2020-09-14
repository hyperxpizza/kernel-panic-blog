package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
)

type NewPost struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}

func GetAllPosts(c *gin.Context) {
	posts, err := database.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &posts)
}

func GetPostByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}

func CreatePost(c *gin.Context) {
	//Extact token from request header
	token := middleware.ExtractToken(c)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authenticated",
		})
		return
	}

	//Extract user's id from token
	claims, ok := middleware.ExtractClaims(token)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while extracting claims",
		})
	}

	userID := fmt.Sprintf("%v", claims["user_id"])

	var post NewPost

	//Parse json into NewPost struct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// Create slug
	postSlug := slug.Make(post.Title)

	// Insert into the database
	//err := database.InsertPost(post.Title, post.Subtitle, post.Content, postSlug, userID)

}

func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}

func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}
