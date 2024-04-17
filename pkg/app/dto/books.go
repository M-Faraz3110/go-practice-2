package dto

import "time"

type CreateBookData struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Count  *int    `json:"count"`
}

type UpdateBookData struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Count  *string `json:"count"`
}

type Book struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Count     int       `json:"count"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
