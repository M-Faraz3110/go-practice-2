package services

import "go-practice/pkg/infra/repos"

type UsersService struct {
	repo *repos.UsersRepository
}

func NewUsersService(repo *repos.UsersRepository) *UsersService {
	return &UsersService{repo: repo}
}
