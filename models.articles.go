package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// demo article list in memory
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// getAllArticles returns a list of all articles in the store
func getAllArticles() []article {
	return articleList
}

// getArticleByID gets a single article given an ID
func getArticleByID(id int) (*article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}
