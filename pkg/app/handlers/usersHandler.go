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

type UsersHandler struct {
	svc *services.UsersService
}

func NewUsersHandler(svc *services.UsersService) *UsersHandler {
	return &UsersHandler{svc: svc}
}

func (hndl *UsersHandler) CreateUsersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//var resp []byte
	switch r.Method {
	case http.MethodPost:
		{
			req := dto.CreateUserRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err) //change to JSON
				return
			}
			res, err := hndl.svc.CreateUser(context.TODO(), req.Email)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
			resp, err := json.Marshal(dto.User{
				Id:        res.Id,
				Email:     res.Email,
				Deleted:   res.Deleted,
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
		w.WriteHeader(405)
		fmt.Fprintln(w, "Method not allowed, only POST allowed for this endpoint")
	}

}

func (hndl *UsersHandler) UsersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var resp []byte
	switch r.Method {
	case http.MethodDelete:
		{
			req := strings.TrimPrefix(r.URL.Path, "/users/")
			res, err := hndl.svc.DeleteUser(context.TODO(), req) //delete borrows as well
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}

			resp, err = json.Marshal(dto.User{
				Id:        res.Id,
				Email:     res.Email,
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
	case http.MethodGet:
		{

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
