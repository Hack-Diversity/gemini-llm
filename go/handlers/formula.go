package handlers

import (
	"llm-gemini/gemini"

	"github.com/gin-gonic/gin"
)

func BuildFormula(c *gin.Context) {
	var input struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	res, err := gemini.BuildFormula(input.Prompt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"formula": res,
	})
}
