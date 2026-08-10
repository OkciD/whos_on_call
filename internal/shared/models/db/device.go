package db

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/shared/errors"
	appModels "github.com/OkciD/whos_on_call/internal/shared/models"
)

type Device struct {
	ID     int
	Name   string
	Type   int8
	UserID int
}

func (d *Device) ToAppModel() (*appModels.Device, error) {
	appDevice := &appModels.Device{
		ID:   d.ID,
		Name: d.Name,
	}

	switch d.Type {
	case int8(appModels.DeviceTypePC):
		fallthrough
	case int8(appModels.DeviceTypeLaptop):
		fallthrough
	case int8(appModels.DeviceTypeMobile):
		appDevice.Type = appModels.DeviceType(d.Type)
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return appDevice, nil
}

func FromDeviceAppModel(appDevice *appModels.Device) (*Device, error) {
	dbDevice := &Device{
		ID:     appDevice.ID,
		Name:   appDevice.Name,
		UserID: appDevice.User.ID,
	}

	switch appDevice.Type {
	case appModels.DeviceTypePC:
		fallthrough
	case appModels.DeviceTypeLaptop:
		fallthrough
	case appModels.DeviceTypeMobile:
		dbDevice.Type = int8(appDevice.Type)
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return dbDevice, nil
}

type DeviceListParams struct {
	Name sql.NullString
	Type sql.NullInt16
}

func FromDeviceListParamsAppModel(appParams *appModels.DeviceListParams) (*DeviceListParams, error) {
	dbParams := &DeviceListParams{}

	if appParams.Name != nil {
		dbParams.Name = sql.NullString{
			String: *appParams.Name,
			Valid:  true,
		}
	}

	if appParams.Type != nil {
		switch *appParams.Type {
		case appModels.DeviceTypePC:
			fallthrough
		case appModels.DeviceTypeLaptop:
			fallthrough
		case appModels.DeviceTypeMobile:
			dbParams.Type = sql.NullInt16{
				Int16: int16(*appParams.Type),
				Valid: true,
			}
		default:
			return nil, errors.ErrDeviceTypeInvalid
		}
	}

	return dbParams, nil
}
