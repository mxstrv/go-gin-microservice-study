package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine // Explicit?

func main() {
	// Set router
	router = gin.Default()
	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Route '/'
	router.GET("/", showIndexPage)
	// Route /article/view/{article_id} handler
	router.GET("/article/view/:article_id", getArticle)

	// Start app
	_ = router.Run()
}
