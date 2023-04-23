package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getAllDocuments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, MockDocuments)
}

func newDocument(c *gin.Context) {
	var d Document

	err := c.BindJSON(&d) // add the json to d
	if err != nil {       // in case there is an error (if it's nil it means there is no error)
		return // return error code made by BindJSON
	}
	if d.CreatedAt.IsZero() {
		d.CreatedAt = time.Now()
	}
	if d.LastEditAt.IsZero() {
		d.LastEditAt = time.Now()
	}
	MockDocuments = append(MockDocuments, d)
	c.IndentedJSON(http.StatusCreated, d)
}

func getDocumentById(id string) (*Document, error) {
	for i, d := range MockDocuments {
		if d.ID == id {
			return &MockDocuments[i], nil
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

func editDocument(c *gin.Context) {
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
	d.LastEditAt = time.Now()

	c.IndentedJSON(http.StatusOK, d)
}
