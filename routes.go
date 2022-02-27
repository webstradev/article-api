package main

func initializeRoutes() {
	// All Articles Route
	router.GET("/", showIndexPage)

	// Single Article View Route
	router.GET("/article/view/:article_id", getArticle)
}
