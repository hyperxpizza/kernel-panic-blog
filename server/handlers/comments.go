package handlers

import (
	"net/http"
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
)

type NewComment struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	IsAdmin bool   `json:"is_admin"`
	OpEmail string `json:"op_email"`
	OpName  string `json:"op_name"`
}

func AddComment(c *gin.Context) {
	postID := c.Param("id")
	postIDinteger, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	var newComment NewComment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	//validate post id
	if postIDinteger != newComment.PostID {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "ids do not match",
		})

		return
	}

	//validate comment
	if len(newComment.Content) > 1000 || len(newComment.Content) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Comment empty or too long",
		})
		return
	}

	//validate op_email
	err = checkmail.ValidateFormat(newComment.OpEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err = checkmail.ValidateHost(newComment.OpEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if len(newComment.OpEmail) > 200 || len(newComment.OpEmail) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "opEmail empty or too long",
		})
		return
	}

	//validate op_name
	if len(newComment.OpName) > 200 || len(newComment.OpName) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "opEmail empty or too long",
		})
		return
	}

	err = database.InsertComment(newComment.PostID, newComment.IsAdmin, newComment.Content, newComment.OpEmail, newComment.OpName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Comment successfully added",
	})

}

func GetComments(c *gin.Context) {
	postID := c.Param("id")
	postIDinteger, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	comments, err := database.GetCommentsByPostID(postIDinteger)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"comments": &comments,
	})

}
