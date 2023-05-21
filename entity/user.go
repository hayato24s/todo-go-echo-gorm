package entity

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Password string
}

func (u *User) Validate() error {
	if u.ID == uuid.Nil {
		return apperr.ErrValidation
	}

	if err := u.ValidateName(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateName() error {
	if ok := regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(u.Name); !ok {
		return fmt.Errorf("%w : Only alphanumeric characters are allowed in the name", apperr.ErrValidation)
	}
	if length := utf8.RuneCountInString(u.Name); length < 4 || 20 < length {
		return fmt.Errorf("%w : Name must be at least 4 and no more than 20 characters long", apperr.ErrValidation)
	}
	return nil
}

func (u *User) ValidatePassword() error {
	if ok := regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(u.Password); !ok {
		return fmt.Errorf("%w : Only alphanumeric characters are allowed in the password", apperr.ErrValidation)
	}
	if length := utf8.RuneCountInString(u.Password); length < 8 || 20 < length {
		return fmt.Errorf("%w : Password must be at least 8 and no more than 20 characters long", apperr.ErrValidation)
	}
	return nil
}
