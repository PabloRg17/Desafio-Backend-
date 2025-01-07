package main

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Synopsis string `json:"synopsis"`
}
