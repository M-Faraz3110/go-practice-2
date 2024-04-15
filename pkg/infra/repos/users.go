package repos

import (
	"context"
	"database/sql"
	"go-practice/pkg/domain/domains/users"
	"go-practice/pkg/infra/db"
	"time"

	uuid "github.com/satori/go.uuid"
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
	id := uuid.NewV4()
	time := time.Now().UTC()
	res := u.db.QueryRowContext(ctx, insertUserQuery, id, email, time, time) //?... create a trigger later
	// if err != nil {
	// 	return users.User{}, err
	// }
	usr := db.User{}
	err := res.Scan(&usr.Id, &usr.Email, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return users.User{}, err
	}
	return users.User{
		Id:        usr.Id,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}, nil
}

// Delete implements users.IRepository.
func (u *UsersRepository) Delete(ctx context.Context, id string) (users.User, error) {
	res := u.db.QueryRowContext(ctx, deleteUserQuery, id)
	usr := db.User{}
	err := res.Scan(&usr.Id, &usr.Email, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return users.User{}, err
	}
	return users.User{
		Id:        usr.Id,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}, nil
}

// Update implements users.IRepository.
func (u *UsersRepository) Update(ctx context.Context, id string, email string) (users.User, error) {
	panic("unimplemented")
}

const (
	insertUserQuery = `
	INSERT INTO users (
		id, 
		email,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	deleteUserQuery = `
	UPDATE users 
		SET deleted = true
	WHERE id = $1
	RETURNING *
	`
)
