package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
)

type NewPost struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Content     string `json:"content"`
	LangVersion string `json:"lang"`
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

func GetPostBySlug(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}

func CreatePost(c *gin.Context) {
	var newPost NewPost

	//Parse json into NewPost struct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// Create slug
	postSlug := slug.Make(newPost.Title)

	// Check if slug already exists in the database
	if database.CheckIfSlugExists(postSlug) == true {

	}

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

	// convert user's id from string to uuid type
	userIDString := fmt.Sprintf("%v", claims["user_id"])
	userID, err := uuid.FromString(userIDString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// Insert into the database
	post, err := database.InsertPost(newPost.Title, newPost.Subtitle, newPost.Content, postSlug, newPost.LangVersion, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &post)
	return

}

func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uuid type
	postID, err := uuid.FromString(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while converting string to uuid",
		})

		return
	}

	//Check if post exists at all
	postExists := database.CheckIfPostExists(postID)
	if postExists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post with given id does not exits",
		})

		return
	}

	/* Delete the post
	deleted, err := database.DeletePostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}
	*/
}

func GetPostWithLang(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})

	return
}
