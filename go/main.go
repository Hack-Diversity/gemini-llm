package main

import (
	"fmt"
	"llm-gemini/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "^_^")
	})

	router.POST(("/filter"), handlers.FilterJobPosts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + fmt.Sprint(port))
}
