package repos

import (
	"context"
	"database/sql"
	"fmt"
	"go-practice/pkg/domain/domains/books"
	"go-practice/pkg/infra/db"
	"time"

	uuid "github.com/satori/go.uuid"
)

type BooksRepository struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

var _ books.IRepository = (*BooksRepository)(nil)

// Create implements books.IRepository.
func (b *BooksRepository) Create(ctx context.Context, title string, author string, count int) (books.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return books.Book{}, err
	}
	id := uuid.NewV4()
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, insertBookQuery, id, title, author, count, time, time)
	book := db.Book{}
	err = res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return books.Book{}, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return books.Book{}, err
	}
	return books.Book{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		Count:     book.Count,
		Deleted:   book.Deleted,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}, nil
}

// Delete implements books.IRepository.
func (b *BooksRepository) Delete(ctx context.Context, id string) (books.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return books.Book{}, err
	}
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, deleteBookQuery, id, time) //check if all borrows are returned
	book := db.Book{}
	err = res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return books.Book{}, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return books.Book{}, err
	}
	return books.Book{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		Count:     book.Count,
		Deleted:   book.Deleted,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}, nil
}

// Get implements books.IRepository.
func (b *BooksRepository) Get(ctx context.Context, id string) (books.Book, error) {
	res := b.db.QueryRowContext(ctx, getBookQuery, id)
	book := db.Book{}
	err := res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return books.Book{}, err
	}
	return books.Book{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		Count:     book.Count,
		Deleted:   book.Deleted,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}, nil
}

// Update implements books.IRepository.
func (b *BooksRepository) Update(ctx context.Context, id string, title string, auhtor string, count int) (books.Book, error) {
	panic("unimplemented")
}

// BorrowBook implements books.IRepository.
func (b *BooksRepository) BorrowBook(ctx context.Context, bookId string) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res := tx.QueryRowContext(ctx, getBookQuery, bookId) //lock record
	book := db.Book{}
	err = res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	if book.Count > 0 {
		time := time.Now().UTC()
		res := tx.QueryRowContext(ctx, borrowBookQuery, bookId, time)
		err = res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		tx.Rollback()
		return fmt.Errorf("no more copies of this book are left")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

const (
	getBookQuery = `
	SELECT * FROM books
	WHERE id = $1 
	AND NOT deleted
	FOR UPDATE
	`

	borrowBookQuery = `
	UPDATE books
	SET count = count - 1, updated_at = $2 
	WHERE id = $1
	RETURNING *
	`

	insertBookQuery = `
	INSERT INTO books(
		id,
		title,
		author,
		count,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`

	deleteBookQuery = `
	UPDATE books
		SET deleted = true, updated_at = $2
	WHERE id = $1
	RETURNING *
	`
)
