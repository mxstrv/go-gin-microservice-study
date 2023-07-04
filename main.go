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

	// Start app
	router.Run()
}
