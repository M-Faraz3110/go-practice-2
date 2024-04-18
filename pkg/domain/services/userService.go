package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/borrows"
	"go-practice/pkg/domain/domains/users"
	"go-practice/pkg/infra/rmq"
)

type UsersService struct {
	urepo      users.IRepository
	brrepo     borrows.IRepository
	msgChannel *rmq.RMQConnection
}

func NewUsersService(
	urepo users.IRepository,
	brrepo borrows.IRepository,
	msgChannel *rmq.RMQConnection,
) *UsersService {
	return &UsersService{urepo: urepo, brrepo: brrepo, msgChannel: msgChannel}
}

func (svc *UsersService) CreateUser(
	ctx context.Context,
	email string,
) (users.User, error) {
	res, err := svc.urepo.Create(ctx, email)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}
	err = svc.msgChannel.SendMessage(ctx, res, "users", "created")
	if err != nil {
		fmt.Println("failed to send message: ", err)
	}
	return res, nil
}

func (svc *UsersService) DeleteUser(
	ctx context.Context,
	id string,
) (users.User, error) {
	//check if no unreturned borrows
	count, err := svc.brrepo.CountBorrowsByUser(ctx, id)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}

	if count != 0 {
		return users.User{}, fmt.Errorf("this user has yet to return all borrowed books")
	}

	res, err := svc.urepo.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}
	err = svc.msgChannel.SendMessage(ctx, res, "users", "deleted")
	if err != nil {
		fmt.Println("failed to send message: ", err)
	}
	return res, nil
}
