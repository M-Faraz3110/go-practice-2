package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/borrows"
	"go-practice/pkg/domain/domains/users"
)

type UsersService struct {
	urepo  users.IRepository
	brrepo borrows.IRepository
}

func NewUsersService(urepo users.IRepository, brrepo borrows.IRepository) *UsersService {
	return &UsersService{urepo: urepo, brrepo: brrepo}
}

func (svc *UsersService) CreateUser(ctx context.Context, email string) (users.User, error) {
	res, err := svc.urepo.Create(ctx, email)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}

	return res, nil
}

func (svc *UsersService) DeleteUser(ctx context.Context, id string) (users.User, error) {
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

	return res, nil
}
