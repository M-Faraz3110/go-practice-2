package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/books"
	"go-practice/pkg/domain/domains/borrows"
)

type BooksService struct {
	brepo  books.IRepository
	brrepo borrows.IRepository
}

func NewBooksService(brepo books.IRepository, brrepo borrows.IRepository) *BooksService {
	return &BooksService{
		brepo:  brepo,
		brrepo: brrepo,
	}
}

func (svc *BooksService) CreateBook(ctx context.Context, title string, author string, count int) (books.Book, error) {
	res, err := svc.brepo.Create(ctx, title, author, count)
	if err != nil {
		fmt.Println(err)
		return books.Book{}, err
	}

	return res, nil
}

func (svc *BooksService) DeleteBook(ctx context.Context, id string) (books.Book, error) {
	//check if no unreturned borrows
	count, err := svc.brrepo.CountBorrowsByBook(ctx, id)
	if err != nil {
		fmt.Println(err)
		return books.Book{}, err
	}

	if count != 0 {
		return books.Book{}, fmt.Errorf("all copies of this book have yet to be returned")
	}

	res, err := svc.brepo.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return books.Book{}, err
	}

	return res, nil
}
