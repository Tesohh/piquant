package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// get the db and save it in the package level var db
var db = func() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("piquant.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}()

var apikey = func() string {
	key, err := os.ReadFile("apikey.txt")
	if err != nil {
		panic("API Key not set! Please create a `apikey.txt` file with the api key.")
	}
	return string(key)
}()

func main() {
	fmt.Println("API key is", apikey)
	db.AutoMigrate(&Document{})

	router := gin.Default()
	defer router.Run("localhost:8080")
	// Routes
	router.GET("/api/:key/docs", handleGetAllDocuments)
	router.GET("/api/:key/docs/:id", handleGetOneDocument)
	router.POST("/api/:key/new", handleNewDocument)
	router.PATCH("/api/:key/edit/:id", editDocument)
}
