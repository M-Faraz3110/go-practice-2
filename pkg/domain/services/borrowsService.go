package services

import "go-practice/pkg/infra/repos"

type BorrowsService struct {
	repo *repos.BorrowsRepository
}

func NewBorrowsService(repo *repos.BorrowsRepository) *BorrowsService {
	return &BorrowsService{repo: repo}
}
