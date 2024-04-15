package db

import "time"

type User struct {
	Id        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
