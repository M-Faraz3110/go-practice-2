package handlers

import "go-practice/pkg/domain/services"

type BorrowsHandler struct {
	svc *services.BorrowsService
}

func NewBorrowsHandler(svc *services.BorrowsService) *BorrowsHandler {
	return &BorrowsHandler{svc: svc}
}
