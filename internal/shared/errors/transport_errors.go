package errors

import "errors"

var (
	ErrInternal        = errors.New("internal")
	ErrUnknown         = errors.New("unknown")
	ErrNotFound        = errors.New("not found")
	ErrNotImplemented  = errors.New("not implemented")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrBadJSON         = errors.New("invalid json")
	ErrInvalidUrlParam = errors.New("invalid url param")
)
