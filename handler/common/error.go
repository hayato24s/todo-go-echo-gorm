package common

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/apperr"
)

type ErrorRes struct {
	Message string `json:"message"`
}

var ErrToCode = map[error]int{
	apperr.ErrTaskNotFound:          http.StatusNotFound,
	apperr.ErrUnauthorized:          http.StatusUnauthorized,
	apperr.ErrUserNameAlreadyExists: http.StatusBadRequest,
	apperr.ErrUserNotFound:          http.StatusNotFound,
}
