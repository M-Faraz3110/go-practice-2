package books

import "context"

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
		title string,
		author string,
		count int,
	) (Book, error)

	Delete(
		ctx context.Context,
		id string,
	) (Book, error)

	Get(
		ctx context.Context,
		id string,
	) (Book, error)

	BorrowBook(
		ctx context.Context,
		bookId string,
	) error
}
