package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllDocuments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, MockDocuments)
}

func newDocument(c *gin.Context) {
	var newDocument Document

	err := c.BindJSON(&newDocument) // add the json to newDocument
	if err != nil {                 // in case there is an error (if it's nil it means there is no error)
		return // return error code made by BindJSON
	}
	MockDocuments = append(MockDocuments, newDocument)
	c.IndentedJSON(http.StatusCreated, newDocument)
}

func getDocumentById(id string) (*Document, error) {
	for _, d := range MockDocuments {
		if d.ID == id {
			return &d, nil
		}
	}
	return nil, errors.New("no document found")
}

func getOneDocument(c *gin.Context) {
	id := c.Param("id")
	d, err := getDocumentById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Document not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, d)
}
