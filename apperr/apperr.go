package apperr

import "errors"

// basic

var ErrValidation = errors.New("validation failed")
var ErrUnauthorized = errors.New("unauthorized")

// task

var ErrTaskNotFound = errors.New("task not found")

// user

var ErrUserNameAlreadyExists = errors.New("user name already exists")
var ErrUserNotFound = errors.New("user not found")
