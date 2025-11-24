package errors

import "errors"

var (
	Unauthorized    = errors.New("unauthorized")
	ErrUserNotFound = errors.New("user not found")
)
