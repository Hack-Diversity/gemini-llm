package main

import (
	"fmt"
	"llm-gemini/handlers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "^_^")
	})

	router.POST(("/filter"), handlers.FilterJobPosts)
	router.POST(("/formula"), handlers.BuildFormula)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + fmt.Sprint(port))
}
