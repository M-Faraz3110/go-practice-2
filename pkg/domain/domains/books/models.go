package books

import "time"

type Book struct {
	Id        string
	Title     string
	Author    string
	Count     int
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
