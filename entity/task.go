package entity

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
)

type Task struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Completed bool
	CreatedAt time.Time
}

func (t *Task) Validate() error {
	if t.ID == uuid.Nil {
		return apperr.ErrValidation
	}

	if t.UserID == uuid.Nil {
		return apperr.ErrValidation
	}

	if err := t.ValidateTitle(); err != nil {
		return err
	}

	if t.CreatedAt.IsZero() {
		return fmt.Errorf("%v : CreatedAt must not be zero value", apperr.ErrValidation)
	}

	return nil
}

func (t *Task) ValidateTitle() error {
	if length := utf8.RuneCountInString(t.Title); length < 1 || 20 < length {
		return fmt.Errorf("%v : Title must be at least 1 and no more than 20 characters long", apperr.ErrValidation)
	}
	return nil
}
