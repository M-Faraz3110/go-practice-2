package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/books"
	"go-practice/pkg/domain/domains/borrows"
	"go-practice/pkg/domain/domains/users"
	"go-practice/pkg/domain/integrations"
	"go-practice/pkg/infra/rmq"
)

type BorrowsService struct {
	urepo      users.IRepository
	brepo      books.IRepository
	brrepo     borrows.IRepository
	txHandler  integrations.ITransaction
	msgChannel *rmq.RMQConnection
}

func NewBorrowsService(urepo users.IRepository, brepo books.IRepository, brrepo borrows.IRepository, txHandler integrations.ITransaction, msgChannel *rmq.RMQConnection) *BorrowsService {
	return &BorrowsService{
		urepo:      urepo,
		brepo:      brepo,
		brrepo:     brrepo,
		txHandler:  txHandler,
		msgChannel: msgChannel,
	}
}

func (svc *BorrowsService) CreateBorrow(ctx context.Context, userId string, bookId string) (borrows.Borrow, error) {
	tx, err := svc.txHandler.BeginTx(ctx)
	defer svc.txHandler.RollbackTx(tx) //cant roll back a committed transaction anyway
	if err != nil {
		return borrows.Borrow{}, err
	}

	user, err := svc.urepo.Get(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	if user.Id == "" {
		return borrows.Borrow{}, fmt.Errorf("user does not exist")
	}

	book, err := svc.brepo.Get(ctx, bookId, tx) //record locked
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	if book.Id == "" {
		return borrows.Borrow{}, fmt.Errorf("book does not exist")
	}
	if book.Count < 1 {
		return borrows.Borrow{}, fmt.Errorf("no more copies of this book left")
	}

	res, err := svc.brrepo.Borrow(ctx, bookId, userId, tx)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}

	err = svc.brepo.BorrowBook(ctx, bookId, tx)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}

	err = svc.txHandler.CommitTx(tx)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, err
	}
	//maybe have the transaction here?
	err = svc.msgChannel.SendMessage(ctx, res, "borrows", "created")
	if err != nil {
		fmt.Println("failed to send message: ", err)
	}
	return res, nil
}

func (svc *BorrowsService) ReturnBorrow(ctx context.Context, borrowId string) (borrows.Borrow, error) {
	res, err := svc.brrepo.Returned(ctx, borrowId)
	if err != nil {
		fmt.Println(err)
		return borrows.Borrow{}, nil
	}
	err = svc.msgChannel.SendMessage(ctx, res, "borrows", "deleted")
	if err != nil {
		fmt.Println("failed to send message: ", err)
	}
	return res, nil
}
