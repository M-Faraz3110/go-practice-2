package db

import "time"

type User struct {
	Id        string
	Email     string
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
