package dto

import "time"

type CreateBorrowData struct {
	UserId *string `json:"userId"`
	BookId *string `json:"bookId"`
}

type Borrow struct {
	Id        string
	BookId    string
	UserId    string
	Returned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
