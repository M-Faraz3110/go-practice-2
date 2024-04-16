package services

import (
	"context"
	"fmt"
	"go-practice/pkg/domain/domains/users"
	"go-practice/pkg/infra/repos"
)

type UsersService struct {
	repo *repos.UsersRepository
}

func NewUsersService(repo *repos.UsersRepository) *UsersService {
	return &UsersService{repo: repo}
}

func (svc *UsersService) CreateUser(ctx context.Context, email string) (users.User, error) {
	res, err := svc.repo.Create(ctx, email)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}

	return res, nil
}

func (svc *UsersService) DeleteUser(ctx context.Context, Id string) (users.User, error) {
	res, err := svc.repo.Delete(ctx, Id)
	if err != nil {
		fmt.Println(err)
		return users.User{}, err
	}

	return res, nil
}
