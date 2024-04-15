package books

import "time"

type Book struct {
	Id     string
	Title  string
	Author string
	Date   time.Time
	Count  int
}

type CreateBookData struct {
	Title  string
	Author string
	Count  string
}

type UpdateBookData struct {
	Title  *string
	Author *string
	Count  *string
}
