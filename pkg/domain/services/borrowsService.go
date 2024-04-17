package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/books"
	"go-practice/pkg/domain/domains/borrows"
	"go-practice/pkg/domain/domains/users"
)

type BorrowsService struct {
	urepo  users.IRepository
	brepo  books.IRepository
	brrepo borrows.IRepository
}

func NewBorrowsService(urepo users.IRepository, brepo books.IRepository, brrepo borrows.IRepository) *BorrowsService {
	return &BorrowsService{
		urepo:  urepo,
		brepo:  brepo,
		brrepo: brrepo,
	}
}

func (svc *BorrowsService) CreateBorrow(ctx context.Context, userId string, bookId string) (borrows.Borrow, error) {

	user, err := svc.urepo.Get(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	if user.Id == "" {
		return borrows.Borrow{}, fmt.Errorf("user does not exist")
	}
	book, err := svc.brepo.Get(ctx, bookId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	if book.Id == "" {
		return borrows.Borrow{}, fmt.Errorf("book does not exist")
	}

	res, err := svc.brrepo.Borrow(ctx, bookId, userId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}

	err = svc.brepo.BorrowBook(ctx, bookId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	//maybe have the transaction here?
	return res, nil
}

func (svc *BorrowsService) ReturnBorrow(ctx context.Context, borrowId string) (borrows.Borrow, error) {
	res, err := svc.brrepo.Returned(ctx, borrowId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, nil
	}

	return res, nil
}
