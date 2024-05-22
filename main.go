package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)
	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)
	router.Run()
}
