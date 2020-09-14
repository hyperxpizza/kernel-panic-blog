package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type NewPost struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
}

func GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
}

func GetPostByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
}

func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
	return
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
