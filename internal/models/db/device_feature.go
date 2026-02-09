package db

import (
	"database/sql"

	appModels "github.com/OkciD/whos_on_call/internal/models"
	"github.com/OkciD/whos_on_call/internal/server/pkg/errors"
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

	if appDeviceFeature.LastActive != nil {
		dbDeviceFeature.LastActive = sql.NullTime{
			Valid: true,
			Time:  *appDeviceFeature.LastActive,
		}
	}

	return dbDeviceFeature, nil
}
