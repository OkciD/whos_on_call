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

func (p *ListDevicesParams) ToAppModel() (*appModels.DeviceListParams, error) {
	appParams := &appModels.DeviceListParams{}

	if p.Type != nil {
		switch *p.Type {
		case DeviceTypeLaptop:
			appParams.Type = new(appModels.DeviceTypeLaptop)
		case DeviceTypeMobile:
			appParams.Type = new(appModels.DeviceTypeMobile)
		case DeviceTypePC:
			appParams.Type = new(appModels.DeviceTypePC)
		default:
			return nil, errors.ErrDeviceTypeInvalid
		}
	}

	if p.Name != nil {
		appParams.Name = p.Name
	}

	if appParams.Name == nil && appParams.Type == nil {
		return nil, nil
	}

	return appParams, nil
}

func FromDeviceListParamsAppModel(appParams *appModels.DeviceListParams) (*ListDevicesParams, error) {
	apiParams := &ListDevicesParams{}

	if appParams.Type != nil {
		switch *appParams.Type {
		case appModels.DeviceTypeLaptop:
			apiParams.Type = new(DeviceTypeLaptop)
		case appModels.DeviceTypeMobile:
			apiParams.Type = new(DeviceTypeMobile)
		case appModels.DeviceTypePC:
			apiParams.Type = new(DeviceTypePC)
		default:
			return nil, errors.ErrDeviceTypeInvalid
		}
	}

	if appParams.Name != nil {
		apiParams.Name = appParams.Name
	}

	if apiParams.Name == nil && apiParams.Type == nil {
		return nil, nil
	}

	return apiParams, nil
}
