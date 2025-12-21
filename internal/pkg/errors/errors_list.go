package errors

import "errors"

// todo: разделить ошибки приложения и транспортного уровня

var (
	ErrInternal       = errors.New("internal")
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemented")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidJSON    = errors.New("invalid json")

	ErrUserNotFound = errors.New("user not found")

	ErrDeviceTypeInvalid = errors.New("invalid device type")
	ErrDeviceExists      = errors.New("device already exists")

	ErrDeviceFeatureTypeInvalid   = errors.New("invalid device feature type")
	ErrDeviceFeatureStatusInvalid = errors.New("invalid device feature status")
)
