package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/entity"
)

type IUserRepository interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *entity.User) error

	// FindUserByName returns the user specified by Name.
	//
	// The following errors may be returned.
	//   - ErrUserNotFound
	FindUserByName(ctx context.Context, name string) (*entity.User, error)

	// Delete the user specified by the given id.
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}
