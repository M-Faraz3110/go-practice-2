package services

import "go-practice/pkg/infra/repos"

type BooksService struct {
	repo *repos.BooksRepository
}

func NewBooksService(repo *repos.BooksRepository) *BooksService {
	return &BooksService{repo: repo}
}
