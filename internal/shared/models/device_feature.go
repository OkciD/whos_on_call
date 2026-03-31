package models

import (
	"strconv"
	"time"
)

type DeviceFeatureType int8

const (
	DeviceFeatureTypeMic DeviceFeatureType = iota
	DeviceFeatureTypeCamera
)

var humanReadableDeviceFeatureType = map[DeviceFeatureType]string{
	DeviceFeatureTypeMic:    "mic",
	DeviceFeatureTypeCamera: "camera",
}

func (t DeviceFeatureType) String() string {
	if str, ok := humanReadableDeviceFeatureType[t]; ok {
		return str
	} else {
		return strconv.FormatInt(int64(t), 10)
	}
}

type DeviceFeatureStatus int8

const (
	DeviceFeatureStatusInactive DeviceFeatureStatus = iota
	DeviceFeatureStatusActive
)

var humanReadableDeviceFeatureStatus = map[DeviceFeatureStatus]string{
	DeviceFeatureStatusInactive: "inactive",
	DeviceFeatureStatusActive:   "active",
}

func (t DeviceFeatureStatus) String() string {
	if str, ok := humanReadableDeviceFeatureStatus[t]; ok {
		return str
	} else {
		return strconv.FormatInt(int64(t), 10)
	}
}

type DeviceFeature struct {
	ID         int
	Type       DeviceFeatureType
	Status     DeviceFeatureStatus
	LastActive *time.Time
	Device     *Device
}

func (f *DeviceFeature) WasActiveRecently(timeDelta time.Duration) bool {
	return f.LastActive != nil && time.Since(*f.LastActive) <= timeDelta
}
