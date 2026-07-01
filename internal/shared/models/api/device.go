package api

import (
	"github.com/OkciD/whos_on_call/internal/shared/errors"
	appModels "github.com/OkciD/whos_on_call/internal/shared/models"
)

func (d *Device) ToAppModel() (*appModels.Device, error) {
	appDevice := &appModels.Device{
		ID:   int(d.Id),
		Name: d.Name,
	}

	switch d.Type {
	case DeviceTypeLaptop:
		appDevice.Type = appModels.DeviceTypeLaptop
	case DeviceTypeMobile:
		appDevice.Type = appModels.DeviceTypeMobile
	case DeviceTypePC:
		appDevice.Type = appModels.DeviceTypePC
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return appDevice, nil
}

func FromDeviceAppModel(appDevice *appModels.Device) (*Device, error) {
	apiDevice := &Device{
		Id:   int32(appDevice.ID),
		Name: appDevice.Name,
	}

	switch appDevice.Type {
	case appModels.DeviceTypeLaptop:
		apiDevice.Type = DeviceTypeLaptop
	case appModels.DeviceTypeMobile:
		apiDevice.Type = DeviceTypeMobile
	case appModels.DeviceTypePC:
		apiDevice.Type = DeviceTypePC
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return apiDevice, nil
}
