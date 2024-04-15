package users

import (
	"fmt"
	"go-practice/pkg/domain/services"
	"net/http"
)

type UsersHandler struct {
	svc *services.UsersService
}

func NewUsersHandler(svc *services.UsersService) *UsersHandler {
	return &UsersHandler{svc: svc}
}

func (*UsersHandler) UsersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
}
