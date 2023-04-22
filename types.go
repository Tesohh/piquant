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
	Author     Author    `json:"author"`
	LastEditAt time.Time `json:"lasteditat"`
	CreatedAt  time.Time `json:"createdat"`
}
