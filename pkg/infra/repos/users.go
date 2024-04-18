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
func (u *UsersRepository) Create(
	ctx context.Context,
	email string,
) (users.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return users.User{}, err
	}
	id := uuid.NewV4()
	time := time.Now().UTC()
	res := tx.QueryRowContext(
		ctx, insertUserQuery, id, email, time, time) //?... create a trigger later
	// if err != nil {
	// 	return users.User{}, err
	// }
	usr := db.User{}
	err = res.Scan(
		&usr.Id,
		&usr.Email,
		&usr.Deleted,
		&usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return users.User{}, err
	}
	err = tx.Commit()
	if err != nil {
		return users.User{}, err
	}
	return users.User{
		Id:        usr.Id,
		Email:     usr.Email,
		Deleted:   usr.Deleted,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}, nil
}

// Delete implements users.IRepository.
func (u *UsersRepository) Delete(
	ctx context.Context,
	id string,
) (users.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return users.User{}, err
	}
	time := time.Now().UTC()
	res := tx.QueryRowContext(ctx, deleteUserQuery, id, time) //check if all borrows returned
	usr := db.User{}
	err = res.Scan(
		&usr.Id,
		&usr.Email,
		&usr.Deleted, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return users.User{}, err
	}
	err = tx.Commit()
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
func (u *UsersRepository) Update(
	ctx context.Context,
	id string,
	email string,
) (users.User, error) {
	panic("unimplemented")
}

// Get implements users.IRepository.
func (u *UsersRepository) Get(
	ctx context.Context,
	id string,
) (users.User, error) {
	res := u.db.QueryRowContext(ctx, getUserQuery, id)
	usr := db.User{}
	err := res.Scan(
		&usr.Id,
		&usr.Email,
		&usr.Deleted, &usr.CreatedAt, &usr.UpdatedAt)
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
		SET deleted = true, updated_at = $2
	WHERE id = $1
	RETURNING *
	`

	getUserQuery = `
	SELECT * FROM users
	WHERE id = $1
	FOR UPDATE
	`
)
