package handlers

import "go-practice/pkg/domain/services"

type BooksHandler struct {
	svc *services.BooksService
}

func NewBooksHandler(svc *services.BooksService) *BooksHandler {
	return &BooksHandler{svc: svc}
}
