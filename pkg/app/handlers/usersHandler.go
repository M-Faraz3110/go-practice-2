package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-practice/pkg/app/dto"
	"go-practice/pkg/domain/services"
	"go-practice/pkg/infra/trace"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type UsersHandler struct {
	svc *services.UsersService
}

func NewUsersHandler(svc *services.UsersService) *UsersHandler {
	return &UsersHandler{svc: svc}
}

func (hndl *UsersHandler) CreateUsersHandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	var traceId string
	if traceHeader, ok := r.Header["Trace-Id"]; len(traceHeader) != 1 && !ok {
		fmt.Println("traceId header not formatted correctly")
		traceId = uuid.NewV4().String()
	} else {
		traceId = traceHeader[0]
	}
	//var resp []byte
	switch r.Method {
	case http.MethodPost:
		{
			req := dto.CreateUserData{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err) //change to JSON
				return
			}
			if req.Email == nil {
				w.WriteHeader(400)
				fmt.Println(w, "attributes missing")
				return
			}
			res, err := hndl.svc.CreateUser(
				context.WithValue(context.TODO(), trace.ContextKey("traceId"), traceId),
				*req.Email)
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
		{
			w.WriteHeader(405)
			fmt.Fprintln(w, "Method not allowed, only POST allowed for this endpoint")
		}

	}

}

func (hndl *UsersHandler) UsersHandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	var traceId string
	if traceHeader, ok := r.Header["Trace-Id"]; len(traceHeader) != 1 && !ok {
		fmt.Println("traceId header not formatted correctly")
		traceId = uuid.NewV4().String()
	} else {
		traceId = traceHeader[0]
	}
	var resp []byte
	switch r.Method {
	case http.MethodDelete:
		{
			req := strings.TrimPrefix(r.URL.Path, "/users/")
			res, err := hndl.svc.DeleteUser(
				context.WithValue(context.TODO(), trace.ContextKey("traceId"), traceId),
				req) //delete borrows as well
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
	case http.MethodPatch:
		{
			req := strings.TrimPrefix(r.URL.Path, "/users/")
			reqBody := dto.CreateUserData{}
			err := json.NewDecoder(r.Body).Decode(&reqBody)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err.Error())
				return
			}
			if reqBody.Email == nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, "attributes missing")
				return
			}

			res, err := hndl.svc.UpdateUser(
				context.TODO(),
				req,
				reqBody.Email,
			)
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
