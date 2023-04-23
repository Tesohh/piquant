package main

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	// ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Document struct {
	gorm.Model
	// ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID string `json:"authorid"` // gorm:"embedded"
	Hidden   bool   `json:"hidden"`
}
