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

	c.HTML(
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
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

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   "Article 1",
			"payload": article,
		},
	)
}
