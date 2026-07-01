package api

import (
	"github.com/OkciD/whos_on_call/internal/shared/errors"
	appModels "github.com/OkciD/whos_on_call/internal/shared/models"
)

func (d *DeviceFeature) ToAppModel() (*appModels.DeviceFeature, error) {
	appDeviceFeature := &appModels.DeviceFeature{
		ID: int(d.Id),
	}

	switch d.Type {
	case DeviceFeatureTypeMic:
		appDeviceFeature.Type = appModels.DeviceFeatureTypeMic
	case DeviceFeatureTypeCamera:
		appDeviceFeature.Type = appModels.DeviceFeatureTypeCamera
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch d.Status {
	case DeviceFeatureStatusInactive:
		appDeviceFeature.Status = appModels.DeviceFeatureStatusInactive
	case DeviceFeatureStatusActive:
		appDeviceFeature.Status = appModels.DeviceFeatureStatusActive
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	return appDeviceFeature, nil
}

func FromDeviceFeatureAppModel(appDeviceFeature *appModels.DeviceFeature) (*DeviceFeature, error) {
	apiDeviceFeature := &DeviceFeature{
		Id:         int32(appDeviceFeature.ID),
		LastActive: appDeviceFeature.LastActive,
	}

	switch appDeviceFeature.Type {
	case appModels.DeviceFeatureTypeMic:
		apiDeviceFeature.Type = DeviceFeatureTypeMic
	case appModels.DeviceFeatureTypeCamera:
		apiDeviceFeature.Type = DeviceFeatureTypeCamera
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch appDeviceFeature.Status {
	case appModels.DeviceFeatureStatusInactive:
		apiDeviceFeature.Status = DeviceFeatureStatusInactive
	case appModels.DeviceFeatureStatusActive:
		apiDeviceFeature.Status = DeviceFeatureStatusActive
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	return apiDeviceFeature, nil
}
