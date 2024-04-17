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

type BooksHandler struct {
	svc *services.BooksService
}

func NewBooksHandler(svc *services.BooksService) *BooksHandler {
	return &BooksHandler{svc: svc}
}

func (hndl *BooksHandler) CreateBooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		{
			req := dto.CreateBookData{} //check for missing attributes
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err.Error())
				return
			}
			if req.Author == nil || req.Title == nil || req.Count == nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, "attributes missing")
				return
			}
			res, err := hndl.svc.CreateBook(context.TODO(), *req.Title, *req.Author, *req.Count)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err.Error())
				return
			}
			resp, err := json.Marshal(dto.Book{
				Id:        res.Id,
				Title:     res.Title,
				Author:    res.Author,
				Count:     res.Count,
				Deleted:   res.Deleted,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			})
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err.Error())
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

func (hndl *BooksHandler) BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var resp []byte
	switch r.Method {
	case http.MethodDelete:
		{
			req := strings.TrimPrefix(r.URL.Path, "/books/")
			res, err := hndl.svc.DeleteBook(context.TODO(), req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}

			resp, err = json.Marshal(dto.Book{
				Id:        res.Id,
				Title:     res.Title,
				Author:    res.Author,
				Count:     res.Count,
				Deleted:   res.Deleted,
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
