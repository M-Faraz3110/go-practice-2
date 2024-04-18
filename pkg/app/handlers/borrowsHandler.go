package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-practice/pkg/app/dto"
	"go-practice/pkg/domain/services"
	"net/http"
	"strings"
)

type BorrowsHandler struct {
	svc *services.BorrowsService
}

func NewBorrowsHandler(svc *services.BorrowsService) *BorrowsHandler {
	return &BorrowsHandler{svc: svc}
}

func (hndl *BorrowsHandler) CreateBorrowHandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodPost:
		{
			req := dto.CreateBorrowData{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
			if req.UserId == nil || req.BookId == nil {
				w.WriteHeader(400)
				fmt.Println(w, "attributes missing")
				return
			}
			res, err := hndl.svc.CreateBorrow(context.TODO(), *req.UserId, *req.BookId)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
			resp, err := json.Marshal(dto.Borrow{
				Id:        res.Id,
				BookId:    res.BookId,
				UserId:    res.UserId,
				Returned:  res.Returned,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			})
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
		}
	default:
		{
			w.WriteHeader(405)
			fmt.Fprintln(w, "Method not allowed, only POST allowed for this endpoint")
		}

	}
}

func (hndl *BorrowsHandler) BorrowHandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	var resp []byte
	switch r.Method {
	case http.MethodDelete:
		{
			req := strings.TrimPrefix(r.URL.Path, "/borrows/")
			res, err := hndl.svc.ReturnBorrow(context.TODO(), req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}

			resp, err = json.Marshal(dto.Borrow{
				Id:        res.Id,
				BookId:    res.BookId,
				UserId:    res.UserId,
				Returned:  res.Returned,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			})
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
