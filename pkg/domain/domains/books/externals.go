package books

import (
	"context"
	"database/sql"
)

type IRepository interface {
	Create(
		ctx context.Context,
		title string,
		author string,
		count int,
	) (Book, error)

	Update(
		ctx context.Context,
		id string,
		title *string,
		author *string,
		count *int,
	) (Book, error)

	Delete(
		ctx context.Context,
		id string,
	) (Book, error)

	Get(
		ctx context.Context,
		id string,
		tx *sql.Tx,
	) (Book, error)

	BorrowBook(
		ctx context.Context,
		bookId string,
		tx *sql.Tx,
	) error
}
