package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// show all articles handler
func showIndexPage(c *gin.Context) {
	// Get articles from the store
	articles := getAllArticles()

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

// show single article handler
func getArticle(c *gin.Context) {
	// Check if the article id sucessfully parses to a string
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Check if article exists
	article, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	render(
		c,
		gin.H{
			"title":   "Article 1",
			"payload": article,
		},
		"article.html",
	)
}

// render function based on 'Accept' header
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// JSON Response
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// XML Response
		c.XML(http.StatusOK, data["payload"])
	default:
		// HTML Response with template
		c.HTML(http.StatusOK, templateName, data)
	}
}
