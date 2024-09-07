package handlers

import (
	"llm-gemini/gemini"

	"github.com/gin-gonic/gin"
)

func FilterJobPosts(c *gin.Context) {
	var input struct {
		Posts []string `json:"posts"`
		Title string   `json:"title"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	res, err := gemini.FilterSocialMediaPosts(input.Posts, input.Title)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": res,
	})
}
