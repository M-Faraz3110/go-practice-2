package users

import (
	"context"
	"encoding/json"
	"fmt"
	"go-practice/pkg/app/dto"
	"go-practice/pkg/domain/services"
	"net/http"
)

type UsersHandler struct {
	svc *services.UsersService
}

func NewUsersHandler(svc *services.UsersService) *UsersHandler {
	return &UsersHandler{svc: svc}
}

func (hndl *UsersHandler) CreateUsersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var resp []byte
	switch r.Method {
	case http.MethodPost:
		{
			req := dto.CreateUserRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
			res, err := hndl.svc.CreateUser(context.TODO(), req.Email)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
			resp, err = json.Marshal(dto.User{
				Id:        res.Id,
				Email:     res.Email,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			})
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err)
				return
			}
		}
	default:
		w.WriteHeader(405)
		fmt.Fprintln(w, "Method not allowed, only POST allowed for this endpoint")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}
