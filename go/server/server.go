package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(port int) {
	router := gin.Default()

	// Define the routes
	router.GET("/", healthCheck)
	router.POST(("/filter"), filterJobPosts)

	// Start the server
	router.Run(":" + fmt.Sprint(port))
}
