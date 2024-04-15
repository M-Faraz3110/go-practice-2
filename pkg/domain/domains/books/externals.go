package books

import "context"

type IRepository interface {
	Create(
		ctx context.Context,
		book CreateBookData,
	) (Book, error)

	Update(
		ctx context.Context,
		id string,
		book UpdateBookData,
	) (Book, error)

	Delete(
		ctx context.Context,
		id string,
	) (Book, error)
}
