package errors

import "errors"

var (
	ErrInternal     = errors.New("internal")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrUserNotFound = errors.New("user not found")
)
