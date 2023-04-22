package main

import "time"

type Author struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Document struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	AuthorID   string    `json:"authorid"`
	LastEditAt time.Time `json:"lasteditat"`
	CreatedAt  time.Time `json:"createdat"`
}
