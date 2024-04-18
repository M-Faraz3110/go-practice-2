package borrows

import "time"

type Borrow struct {
	Id        string    `json:"id"`
	BookId    string    `json:"bookId"`
	UserId    string    `json:"userId"`
	Returned  bool      `json:"returned"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
