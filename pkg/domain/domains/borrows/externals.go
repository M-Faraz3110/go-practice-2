package borrows

import "context"

type IRepository interface {
	Borrow(
		ctx context.Context,
		bookId string,
		userId string,
	) (Borrow, error)

	Returned(
		ctx context.Context,
		borrowId string,
	) (Borrow, error)

	CountBorrowsByBook(
		ctx context.Context,
		id string,
	) (int, error)

	CountBorrowsByUser(
		ctx context.Context,
		id string,
	) (int, error)
}
