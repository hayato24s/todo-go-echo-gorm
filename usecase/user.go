package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
)

type CreateUserIn struct {
	Name     string
	Password string
}

// CreateUser creates a new user with the given name and password.
//
// The following errors may be returned.
//   - ErrUserNameAlreadyExists
//   - ErrValidation
func (uc *UseCase) CreateUser(ctx context.Context, in *CreateUserIn) (*entity.User, error) {
	user := &entity.User{
		ID:       uuid.New(),
		Name:     in.Name,
		Password: in.Password,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	_, err := uc.r.FindUserByName(ctx, user.Name)
	if err != nil {
		if errors.Is(err, apperr.ErrUserNotFound) {
			return user, uc.r.CreateUser(ctx, user)
		}
		return nil, err
	}
	return nil, apperr.ErrUserNameAlreadyExists
}

// DeleteUser deletes the user and related resources.
//
// The following errors may be returned.
//   - ErrUnauthorized
func (uc *UseCase) DeleteUser(ctx context.Context) error {
	userID, err := uc.Authenticate(ctx)
	if err != nil {
		return err
	}

	tx, err := uc.r.Begin()
	if err != nil {
		return err
	}

	if err := tx.DeleteTasksByUserID(ctx, userID); err != nil {
		return tx.Rollback()
	}

	if err := tx.DeleteUserByID(ctx, userID); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
