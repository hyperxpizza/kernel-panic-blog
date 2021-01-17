package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
	"github.com/hyperxpizza/kernel-panic-blog/server/middleware"
)

type NewPost struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle"`
	Content  string  `json:"content"`
}

func GetAllPosts(c *gin.Context) {
	posts, err := database.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &posts)

}

func GetPost(c *gin.Context) {
	slug := c.Param("slug")

	/*
		//convert to int
		postIDInteger, err := strconv.Atoi(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	*/

	post, err := database.GetPostBySlug(slug)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"succcess": false,
				"message":  "No post with provided id was found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"post":    &post,
	})

	return
}

func CreatePost(c *gin.Context) {
	var newPost NewPost
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//check if post with given title already exists
	exists := database.CheckIfPostExists(newPost.Title)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Post with given title already exists in the database.",
		})
		return
	}

	token, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	claims, ok := middleware.ExtractClaims(*token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Error while extracting claims",
		})
		return
	}

	userID := claims["user_id"]

	//create slug from title
	slug := slug.Make(newPost.Title)

	err = database.CreatePost(newPost.Title, newPost.Content, slug, int(userID.(float64)), newPost.Subtitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post added successfully",
	})

}

func GetTags(c *gin.Context) {

}

func AddTag(c *gin.Context) {

}
