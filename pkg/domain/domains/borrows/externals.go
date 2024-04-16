package borrows

import "context"

type IRepository interface {
	Borrow(
		ctx context.Context,
		bookId string,
		userId string,
	) (Borrowed, error)

	Returned(
		ctx context.Context,
		borrowId string,
	) (Borrowed, error)
}
