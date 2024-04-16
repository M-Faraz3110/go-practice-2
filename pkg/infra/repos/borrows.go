package repos

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/domains/borrows"
)

type BorrowsRepository struct {
	db *sql.DB
}

func NewBorrowsRepository(db *sql.DB) *BorrowsRepository {
	return &BorrowsRepository{db: db}
}

var _ borrows.IRepository = (*BorrowsRepository)(nil)

// Borrow implements borrows.IRepository.
func (b *BorrowsRepository) Borrow(ctx context.Context, bookId string, userId string) (borrows.Borrowed, error) {
	panic("unimplemented")
}

// Returned implements borrows.IRepository.
func (b *BorrowsRepository) Returned(ctx context.Context, borrowId string) (borrows.Borrowed, error) {
	panic("unimplemented")
}
