package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a fake request
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	// Create a fake request
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test the status code 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Add Accept header so we receive json instead of html
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		pageBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var pageArticles []article
		err = json.Unmarshal(pageBytes, &pageArticles)

		pageOk := err == nil && len(pageArticles) == 2

		return statusOK && pageOk
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/article/view/:article_id", getArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	// Add Accept header so we receive xml instead of html
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		pageBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		var pageArticle article
		err = xml.Unmarshal(pageBytes, &pageArticle)

		pageOk := err == nil && pageArticle.ID == 1 && pageArticle.Title == "Article 1"

		return statusOK && pageOk
	})
}
