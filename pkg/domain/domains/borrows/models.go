package borrows

import "time"

type Borrow struct {
	Id        string
	BookId    string
	UserId    string
	Returned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
