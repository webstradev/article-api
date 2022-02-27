package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Run the remaining tests
	os.Exit(m.Run())
}

// Helper function to create a mocked router used for testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob(
			"templates/*")
	}
	return r
}

// Helper function to process a request and test the response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a recorder which will be used to test the response
	w := httptest.NewRecorder()

	// Serve the router and process the request to the record
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
func saveLists() {
	tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	articleList = tmpArticleList
}
