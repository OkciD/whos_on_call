package models

import "time"

type DeviceFeatureType string

const (
	DeviceFeatureTypeMic    DeviceFeatureType = "mic"
	DeviceFeatureTypeCamera DeviceFeatureType = "camera"
)

type DeviceFeatureStatus string

const (
	DeviceFeatureStatusActive   DeviceFeatureStatus = "active"
	DeviceFeatureStatusInactive DeviceFeatureStatus = "inactive"
)

type DeviceFeature struct {
	Device       *Device
	Type         DeviceFeatureType
	Status       DeviceFeatureStatus
	LastModified time.Time
}
