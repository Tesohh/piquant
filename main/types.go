package main

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	// ID       string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"` // gorm:"embedded"
	Hidden bool   `json:"hidden"`
}
