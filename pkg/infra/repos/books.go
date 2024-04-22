package repos

import (
	"context"
	"database/sql"
	"fmt"
	"go-practice/pkg/domain/domains/books"
	"go-practice/pkg/infra/db"
	"strconv"
	"strings"
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
func (b *BooksRepository) Create(
	ctx context.Context,
	title string,
	author string,
	count int,
) (books.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return books.Book{}, err
	}
	id := uuid.NewV4()
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, insertBookQuery, id, title, author, count, time, time)
	book := db.Book{}
	err = res.Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Count,
		&book.Deleted,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return books.Book{}, err
	}
	err = tx.Commit()
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

// Delete implements books.IRepository.
func (b *BooksRepository) Delete(
	ctx context.Context,
	id string,
) (books.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return books.Book{}, err
	}
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, deleteBookQuery, id, time) //check if all borrows are returned
	book := db.Book{}
	err = res.Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Count,
		&book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return books.Book{}, err
	}
	err = tx.Commit()
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

// Get implements books.IRepository.
func (b *BooksRepository) Get(
	ctx context.Context,
	id string,
	tx *sql.Tx,
) (books.Book, error) {
	res := tx.QueryRowContext(ctx, getBookQuery, id)
	book := db.Book{}
	err := res.Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Count,
		&book.Deleted, &book.CreatedAt, &book.UpdatedAt)

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
func (b *BooksRepository) Update(
	ctx context.Context,
	id string,
	title *string,
	auhtor *string,
	count *int,
) (books.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return books.Book{}, err
	}

	time := time.Now().UTC()
	query := createUpdateQuery(title, auhtor, count)
	res := tx.QueryRowContext(ctx, query, id, time)
	book := db.Book{}
	err = res.Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Count,
		&book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return books.Book{}, err
	}
	err = tx.Commit()
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

func createUpdateQuery(title *string, author *string, count *int) string { //dont wanna use reflections and all that to make this more generic cuz its overkill
	baseQuery := baseUpdateQuery
	var s []string
	if title != nil {
		s = append(s, "title = '"+*title+"'")
	}
	if author != nil {
		s = append(s, "author = '"+*author+"'")
	}
	if count != nil {
		s = append(s, "count = "+strconv.Itoa(*count))
	}
	resQuery := fmt.Sprintf(baseQuery, strings.Join(s, ","))
	return resQuery
}

// BorrowBook implements books.IRepository.
func (b *BooksRepository) BorrowBook(
	ctx context.Context,
	bookId string,
	tx *sql.Tx,
) error {
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, borrowBookQuery, bookId, time)
	book := db.Book{}
	err := res.Scan(&book.Id, &book.Title, &book.Author, &book.Count, &book.Deleted, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

const (
	baseUpdateQuery = `
	UPDATE books
	SET updated_at = $2, %s
	WHERE id = $1
	RETURNING *
	`

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
