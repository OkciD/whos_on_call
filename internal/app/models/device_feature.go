package models

import "time"

type DeviceFeatureType int8

const (
	DeviceFeatureTypeMic DeviceFeatureType = iota
	DeviceFeatureTypeCamera
)

type DeviceFeatureStatus int8

const (
	DeviceFeatureStatusInactive DeviceFeatureStatus = iota
	DeviceFeatureStatusActive
)

type DeviceFeature struct {
	ID           int
	Type         DeviceFeatureType
	Status       DeviceFeatureStatus
	LastModified *time.Time
	// Device       *Device
}
