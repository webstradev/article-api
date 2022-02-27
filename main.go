package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	router := gin.Default()

	// Load the html templates
	router.LoadHTMLGlob("templates/*")

	// Handle the root route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	// Start the server
	router.Run()
}
