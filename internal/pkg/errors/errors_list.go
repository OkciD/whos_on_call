package errors

import "errors"

var (
	ErrInternal       = errors.New("internal")
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemented")
	ErrUnauthorized   = errors.New("unauthorized")

	ErrUserNotFound = errors.New("user not found")

	ErrDeviceTypeInvalid = errors.New("invalid device type")

	ErrDeviceFeatureTypeInvalid   = errors.New("invalid device feature type")
	ErrDeviceFeatureStatusInvalid = errors.New("invalid device feature status")
)
