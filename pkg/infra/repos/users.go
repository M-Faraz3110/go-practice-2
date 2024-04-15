package repos

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/domains/users"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

var _ users.IRepository = (*UsersRepository)(nil)

// Create implements users.IRepository.
func (u *UsersRepository) Create(ctx context.Context, email string) (users.User, error) {
	panic("unimplemented")
}

// Delete implements users.IRepository.
func (u *UsersRepository) Delete(ctx context.Context, id string) (users.User, error) {
	panic("unimplemented")
}

// Update implements users.IRepository.
func (u *UsersRepository) Update(ctx context.Context, id string, email string) (users.User, error) {
	panic("unimplemented")
}
