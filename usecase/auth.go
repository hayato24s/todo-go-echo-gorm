package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/appctx"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
)

// Authenticate authenticates the user and returns the ID.
//
// The following errors may be returned.
//   - ErrUnauthorized
func (uc *UseCase) Authenticate(ctx context.Context) (uuid.UUID, error) {
	userID, ok := appctx.GetUserID(ctx)
	if !ok || userID == uuid.Nil {
		return uuid.Nil, apperr.ErrUnauthorized
	}
	return userID, nil
}

type LogInIn struct {
	Name     string
	Password string
}

func (in *LogInIn) Validate() error {
	u := &entity.User{Name: in.Name, Password: in.Password}
	if err := u.ValidateName(); err != nil {
		return err
	}
	if err := u.ValidatePassword(); err != nil {
		return err
	}
	return nil
}

// LogIn logs in with the given name and password.
//
// The following errors may be returned.
//   - ErrUnauthorized
func (uc *UseCase) LogIn(ctx context.Context, in *LogInIn) (*entity.User, error) {
	given := &entity.User{Name: in.Name, Password: in.Password}
	if err := given.ValidateName(); err != nil {
		return nil, err
	}
	if err := given.ValidatePassword(); err != nil {
		return nil, err
	}

	user, err := uc.r.FindUserByName(ctx, in.Name)
	if err != nil {
		if errors.Is(err, apperr.ErrUserNotFound) {
			return nil, apperr.ErrUnauthorized
		}
		return nil, err
	}

	if user.Password != in.Password {
		return nil, apperr.ErrUnauthorized
	}
	return user, nil
}
