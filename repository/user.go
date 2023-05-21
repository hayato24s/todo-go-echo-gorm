package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"

	"gorm.io/gorm"
)

func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Create(user).
		Error
}

func (r *Repository) FindUserByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		Take(&user).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrUserNotFound
	}
	return &user, nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.User{}).
		Error
}
