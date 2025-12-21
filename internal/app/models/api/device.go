package api

import (
	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
)

type deviceType string

const (
	deviceTypeLaptop deviceType = "laptop"
	deviceTypeMobile deviceType = "mobile"
	deviceTypePC     deviceType = "pc"
)

type Device struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Type deviceType `json:"type"`
}

func (d *Device) ToAppModel() (*appModels.Device, error) {
	appDevice := &appModels.Device{
		ID:   d.ID,
		Name: d.Name,
	}

	switch d.Type {
	case deviceTypeLaptop:
		appDevice.Type = appModels.DeviceTypeLaptop
	case deviceTypeMobile:
		appDevice.Type = appModels.DeviceTypeMobile
	case deviceTypePC:
		appDevice.Type = appModels.DeviceTypePC
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return appDevice, nil
}

func FromDeviceAppModel(appDevice *appModels.Device) (*Device, error) {
	apiDevice := &Device{
		ID:   appDevice.ID,
		Name: appDevice.Name,
	}

	switch appDevice.Type {
	case appModels.DeviceTypeLaptop:
		apiDevice.Type = deviceTypeLaptop
	case appModels.DeviceTypeMobile:
		apiDevice.Type = deviceTypeMobile
	case appModels.DeviceTypePC:
		apiDevice.Type = deviceTypePC
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return apiDevice, nil
}
