package users

import "time"

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
