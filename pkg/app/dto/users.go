package dto

import "time"

type CreateUserRequest struct {
	Email string `json:"email"`
}

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
