package db

import (
	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
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
		Type: appModels.DeviceType(d.Type),
	}

	switch d.Type {
	case int8(appModels.DeviceTypePC):
	case int8(appModels.DeviceTypeLaptop):
	case int8(appModels.DeviceTypeMobile):
		appDevice.Type = appModels.DeviceType(d.Type)
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return appDevice, nil
}

func FromDeviceAppModel(appDevice *appModels.Device, appUser *appModels.User) (*Device, error) {
	dbDevice := &Device{
		ID:     appDevice.ID,
		Name:   appDevice.Name,
		UserID: appUser.ID,
	}

	switch appDevice.Type {
	case appModels.DeviceTypePC:
	case appModels.DeviceTypeLaptop:
	case appModels.DeviceTypeMobile:
		dbDevice.Type = int8(appDevice.Type)
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return dbDevice, nil
}
