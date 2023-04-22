package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	defer router.Run("localhost:8080")
	// Routes
	router.GET("/docs", getAllDocuments)
	router.GET("/docs/:id", getOneDocument)
	router.POST("/new", newDocument)
}
