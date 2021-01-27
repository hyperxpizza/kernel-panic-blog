package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/kernel-panic-blog/server/database"
)

type NewTag struct {
	Tag string `json:"tag"`
}

/*
func GetTags(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	tags, err := database.GetTagsByPostID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"tags":    &tags,
	})

}
*/

func AddTag(c *gin.Context) {
	var newTag NewTag

	if err := c.ShouldBindJSON(&newTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//check if tag already exists in the database
	tagExists := database.CheckIfTagExists(newTag.Tag)
	if tagExists {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Tag already exists in the database",
		})
		return
	}

	err := database.InsertTag(newTag.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Tag successfully added",
	})

	return
}
