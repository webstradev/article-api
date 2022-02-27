package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Create a default gin router
	router = gin.Default()

	// Load the html templates
	router.LoadHTMLGlob("templates/*")

	// Initialize all routes
	initializeRoutes()

	// Start the server
	router.Run()
}
