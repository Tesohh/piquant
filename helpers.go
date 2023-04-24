package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func getDocumentById(id string) (*Document, error) {
	var document Document
	result := db.Model(&Document{}).First(&document, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("no document found")
	}
	return &document, nil
}

func checkApiKey(c *gin.Context) error {
	key := c.Param("key")
	if key != string(apikey) {
		return errors.New("api key is not correct")
	}
	return nil
}
