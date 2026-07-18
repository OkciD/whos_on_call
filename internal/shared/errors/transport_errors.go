package errors

import "errors"

var (
	ErrInternal           = errors.New("internal")
	ErrUnknown            = errors.New("unknown")
	ErrEntityNotFound     = errors.New("entity not found")
	ErrRouteRouteNotFound = errors.New("route not found")
	ErrUnauthorized       = errors.New("unauthorized")
)
