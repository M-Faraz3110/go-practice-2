package repos

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/domains/borrows"
	"go-practice/pkg/infra/db"
	"time"

	uuid "github.com/satori/go.uuid"
)

type BorrowsRepository struct {
	db *sql.DB
}

func NewBorrowsRepository(db *sql.DB) *BorrowsRepository {
	return &BorrowsRepository{db: db}
}

var _ borrows.IRepository = (*BorrowsRepository)(nil)

// Borrow implements borrows.IRepository.
func (b *BorrowsRepository) Borrow(
	ctx context.Context,
	bookId string,
	userId string,
	tx *sql.Tx,
) (borrows.Borrow, error) {
	id := uuid.NewV4()
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, insertBorrowQuery, id, bookId, userId, time, time)
	borrow := db.Borrow{}
	err := res.Scan(
		&borrow.Id,
		&borrow.BookId,
		&borrow.UserId,
		&borrow.Returned,
		&borrow.CreatedAt, &borrow.UpdatedAt)
	if err != nil {
		return borrows.Borrow{}, err
	}
	return borrows.Borrow{
		Id:        borrow.Id,
		BookId:    borrow.BookId,
		UserId:    borrow.UserId,
		Returned:  borrow.Returned,
		CreatedAt: borrow.CreatedAt,
		UpdatedAt: borrow.UpdatedAt,
	}, nil

}

// Returned implements borrows.IRepository.
func (b *BorrowsRepository) Returned(
	ctx context.Context,
	borrowId string,
) (borrows.Borrow, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return borrows.Borrow{}, err
	}
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, returnBorrowQuery, borrowId, time)
	borrow := db.Borrow{}
	err = res.Scan(
		&borrow.Id,
		&borrow.BookId,
		&borrow.UserId,
		&borrow.Returned,
		&borrow.CreatedAt, &borrow.UpdatedAt)
	if err != nil {
		return borrows.Borrow{}, err
	}
	err = tx.Commit()
	if err != nil {
		return borrows.Borrow{}, err
	}
	return borrows.Borrow{
		Id:        borrow.Id,
		BookId:    borrow.BookId,
		UserId:    borrow.UserId,
		Returned:  borrow.Returned,
		CreatedAt: borrow.CreatedAt,
		UpdatedAt: borrow.UpdatedAt,
	}, nil
}

// GetBorrowsByBook implements borrows.IRepository.
func (b *BorrowsRepository) CountBorrowsByBook(
	ctx context.Context,
	id string,
) (int, error) {
	var count int
	res := b.db.QueryRowContext(ctx, getBorrowsByBook, id)
	err := res.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountBorrowsByUser implements borrows.IRepository.
func (b *BorrowsRepository) CountBorrowsByUser(
	ctx context.Context,
	id string,
) (int, error) {
	var count int
	res := b.db.QueryRowContext(ctx, getBorrowsByUser, id)
	err := res.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

const (
	insertBorrowQuery = `
	INSERT INTO borrows(
		id, 
		book_id,
		user_id,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING *
	`

	returnBorrowQuery = `
	UPDATE borrows
		SET returned = true, updated_at = $2
	WHERE id = $1
	RETURNING *
	`

	getBorrowsByBook = `
	SELECT COUNT(id) FROM borrows
	WHERE book_id = $1
	AND returned = false
	`
	getBorrowsByUser = `
	SELECT COUNT(id) FROM borrows
	WHERE user_id = $1
	AND returned = false
	`
)
