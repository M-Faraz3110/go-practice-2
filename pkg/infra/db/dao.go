package db

import "time"

type User struct {
	Id        string
	Email     string
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	Id        string
	Title     string
	Author    string
	Count     int
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Borrow struct {
	Id        string
	BookId    string
	UserId    string
	Returned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
