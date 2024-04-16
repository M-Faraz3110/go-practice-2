package repos

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/domains/books"
)

type BooksRepository struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

var _ books.IRepository = (*BooksRepository)(nil)

// Create implements books.IRepository.
func (b *BooksRepository) Create(ctx context.Context, book books.CreateBookData) (books.Book, error) {
	panic("unimplemented")
}

// Delete implements books.IRepository.
func (b *BooksRepository) Delete(ctx context.Context, id string) (books.Book, error) {
	panic("unimplemented")
}

// Get implements books.IRepository.
func (b *BooksRepository) Get(ctx context.Context, id string) (books.Book, error) {
	panic("unimplemented")
}

// Update implements books.IRepository.
func (b *BooksRepository) Update(ctx context.Context, id string, book books.UpdateBookData) (books.Book, error) {
	panic("unimplemented")
}
