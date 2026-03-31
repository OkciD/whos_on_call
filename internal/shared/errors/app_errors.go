package errors

import "errors"

var (
	ErrDuplicate = errors.New("duplicate")

	ErrDeviceTypeInvalid = errors.New("invalid device type")

	ErrDeviceFeatureNotFound      = errors.New("device feature not found")
	ErrDeviceFeatureTypeInvalid   = errors.New("invalid device feature type")
	ErrDeviceFeatureStatusInvalid = errors.New("invalid device feature status")

	ErrCallStatusStateInvalid = errors.New("call status state invalid")
)
