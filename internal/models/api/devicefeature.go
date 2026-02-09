package api

import (
	"time"

	appModels "github.com/OkciD/whos_on_call/internal/models"
	"github.com/OkciD/whos_on_call/internal/server/pkg/errors"
)

type deviceFeatureType string

const (
	deviceFeatureTypeMic    deviceFeatureType = "mic"
	deviceFeatureTypeCamera deviceFeatureType = "camera"
)

type deviceFeatureStatus string

const (
	deviceFeatureStatusInactive deviceFeatureStatus = "inactive"
	deviceFeatureStatusActive   deviceFeatureStatus = "active"
)

type DeviceFeature struct {
	ID         int                 `json:"id"`
	Type       deviceFeatureType   `json:"type"`
	Status     deviceFeatureStatus `json:"status"`
	LastActive string              `json:"lastActive,omitempty"`
}

func (d *DeviceFeature) ToAppModel() (*appModels.DeviceFeature, error) {
	appDeviceFeature := &appModels.DeviceFeature{
		ID: d.ID,
	}

	switch d.Type {
	case deviceFeatureTypeMic:
		appDeviceFeature.Type = appModels.DeviceFeatureTypeMic
	case deviceFeatureTypeCamera:
		appDeviceFeature.Type = appModels.DeviceFeatureTypeCamera
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch d.Status {
	case deviceFeatureStatusInactive:
		appDeviceFeature.Status = appModels.DeviceFeatureStatusInactive
	case deviceFeatureStatusActive:
		appDeviceFeature.Status = appModels.DeviceFeatureStatusActive
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	return appDeviceFeature, nil
}

func FromDeviceFeatureAppModel(appDeviceFeature *appModels.DeviceFeature) (*DeviceFeature, error) {
	apiDeviceFeature := &DeviceFeature{
		ID: appDeviceFeature.ID,
	}

	switch appDeviceFeature.Type {
	case appModels.DeviceFeatureTypeMic:
		apiDeviceFeature.Type = deviceFeatureTypeMic
	case appModels.DeviceFeatureTypeCamera:
		apiDeviceFeature.Type = deviceFeatureTypeCamera
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch appDeviceFeature.Status {
	case appModels.DeviceFeatureStatusInactive:
		apiDeviceFeature.Status = deviceFeatureStatusInactive
	case appModels.DeviceFeatureStatusActive:
		apiDeviceFeature.Status = deviceFeatureStatusActive
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	if appDeviceFeature.LastActive != nil {
		apiDeviceFeature.LastActive = appDeviceFeature.LastActive.Format(time.RFC3339)
	}

	return apiDeviceFeature, nil
}
