package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetAllDocuments(c *gin.Context) {
	err := checkApiKey(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	var documents []Document
	result := db.Model(&Document{}).Find(&documents)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": result.Error})
		return
	}
	c.IndentedJSON(http.StatusOK, documents)
}

func handleNewDocument(c *gin.Context) {
	{
		err := checkApiKey(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
	}

	var d Document

	err := c.BindJSON(&d) // add the json to d
	if err != nil {       // in case there is an error (if it's nil it means there is no error)
		return // return error code made by BindJSON
	}
	db.Model(&Document{}).Create(&d)
	c.IndentedJSON(http.StatusCreated, d)
}

func handleGetOneDocument(c *gin.Context) {
	{
		err := checkApiKey(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
	}

	id := c.Param("id")
	d, err := getDocumentById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Document not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, d)
}

func editDocument(c *gin.Context) {
	{
		err := checkApiKey(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
	}

	id := c.Param("id")
	d, derr := getDocumentById(id)
	if derr != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Document not found."})
		return
	}
	jerr := c.BindJSON(&d)
	if jerr != nil {
		return
	}
	db.Model(&Document{}).Where("id = ?", id).Updates(d)
	c.IndentedJSON(http.StatusOK, d)
}
