package borrows

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Borrow(
		ctx context.Context,
		bookId string,
		userId string,
		tx *sql.Tx,
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
