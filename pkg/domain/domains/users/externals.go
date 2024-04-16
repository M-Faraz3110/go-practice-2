package users

import "context"

type IRepository interface {
	Get(
		ctx context.Context,
		id string,
	) (User, error)

	Create(
		ctx context.Context,
		email string,
	) (User, error)

	Update(
		ctx context.Context,
		id string,
		email string,
	) (User, error)

	Delete(
		ctx context.Context,
		id string,
	) (User, error)
}
