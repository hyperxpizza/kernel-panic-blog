package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewComment struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	opEmail string `json:"op_email"`
	opName  string `json:"op_name"`
}

func AddComment(c *gin.Context) {
	var newComment NewComment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

}
