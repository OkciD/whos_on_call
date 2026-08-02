package db

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/shared/errors"
	appModels "github.com/OkciD/whos_on_call/internal/shared/models"
)

type DeviceFeature struct {
	ID         int
	Type       int8
	Status     int8
	LastActive sql.NullTime
	DeviceID   int
}

func (df *DeviceFeature) ToAppModel() (*appModels.DeviceFeature, error) {
	appDeviceFeature := &appModels.DeviceFeature{
		ID: df.ID,
	}

	switch df.Type {
	case int8(appModels.DeviceFeatureTypeMic):
		fallthrough
	case int8(appModels.DeviceFeatureTypeCamera):
		appDeviceFeature.Type = appModels.DeviceFeatureType(df.Type)
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch df.Status {
	case int8(appModels.DeviceFeatureStatusInactive):
		fallthrough
	case int8(appModels.DeviceFeatureStatusActive):
		appDeviceFeature.Status = appModels.DeviceFeatureStatus(df.Status)
	}

	if df.LastActive.Valid {
		appDeviceFeature.LastActive = &df.LastActive.Time
	}

	return appDeviceFeature, nil
}

func FromDeviceFeatureAppModel(
	appDeviceFeature *appModels.DeviceFeature,
) (*DeviceFeature, error) {
	dbDeviceFeature := &DeviceFeature{
		ID:       appDeviceFeature.ID,
		DeviceID: appDeviceFeature.Device.ID,
	}

	switch appDeviceFeature.Type {
	case appModels.DeviceFeatureTypeMic:
		fallthrough
	case appModels.DeviceFeatureTypeCamera:
		dbDeviceFeature.Type = int8(appDeviceFeature.Type)
	default:
		return nil, errors.ErrDeviceFeatureTypeInvalid
	}

	switch appDeviceFeature.Status {
	case appModels.DeviceFeatureStatusInactive:
		fallthrough
	case appModels.DeviceFeatureStatusActive:
		dbDeviceFeature.Status = int8(appDeviceFeature.Status)
	default:
		return nil, errors.ErrDeviceFeatureStatusInvalid
	}

	if appDeviceFeature.LastActive != nil {
		dbDeviceFeature.LastActive = sql.NullTime{
			Valid: true,
			Time:  *appDeviceFeature.LastActive,
		}
	}

	return dbDeviceFeature, nil
}
