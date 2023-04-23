package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// get the db and save it in the package level var db
var db = func() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("docs.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}()

func main() {
	db.AutoMigrate(&Document{})

	router := gin.Default()
	defer router.Run("localhost:8080")
	// Routes
	router.GET("/docs", getAllDocuments)
	router.GET("/docs/:id", getOneDocument)
	router.POST("/new", newDocument)
	router.PATCH("/edit/:id", editDocument)
}
