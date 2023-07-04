package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine // Explicit?

func main() {
	// Set router
	router = gin.Default()
	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Route '/'
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start app
	router.Run()
}
