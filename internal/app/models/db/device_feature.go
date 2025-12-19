package db

import (
	"database/sql"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
)

type DeviceFeature struct {
	ID           int
	Type         int8
	Status       int8
	LastModified sql.NullTime
	DeviceID     int
}

func (df *DeviceFeature) ToAppModel() (*appModels.DeviceFeature, error) {
	appDeviceFeature := &appModels.DeviceFeature{
		ID: df.ID,
	}

	switch df.Type {
	case int8(appModels.DeviceFeatureTypeMic):
	case int8(appModels.DeviceFeatureTypeCamera):
		appDeviceFeature.Type = appModels.DeviceFeatureType(df.Type)
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch df.Status {
	case int8(appModels.DeviceFeatureStatusInactive):
	case int8(appModels.DeviceFeatureStatusActive):
		appDeviceFeature.Status = appModels.DeviceFeatureStatus(df.Status)
	}

	if df.LastModified.Valid {
		appDeviceFeature.LastModified = &df.LastModified.Time
	}

	return appDeviceFeature, nil
}

func FromDeviceFeatureAppModel(
	appDeviceFeature *appModels.DeviceFeature,
	appDevice *appModels.Device,
) (*DeviceFeature, error) {
	dbDeviceFeature := &DeviceFeature{
		ID:       appDeviceFeature.ID,
		DeviceID: appDevice.ID,
	}

	switch appDeviceFeature.Type {
	case appModels.DeviceFeatureTypeMic:
	case appModels.DeviceFeatureTypeCamera:
		dbDeviceFeature.Type = int8(appDeviceFeature.Type)
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch appDeviceFeature.Status {
	case appModels.DeviceFeatureStatusInactive:
	case appModels.DeviceFeatureStatusActive:
		dbDeviceFeature.Status = int8(appDeviceFeature.Status)
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	if appDeviceFeature.LastModified != nil {
		dbDeviceFeature.LastModified = sql.NullTime{
			Valid: true,
			Time:  *appDeviceFeature.LastModified,
		}
	}

	return dbDeviceFeature, nil
}
