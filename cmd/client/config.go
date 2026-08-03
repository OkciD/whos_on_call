package main

import (
	"github.com/OkciD/whos_on_call/internal/client/apiclient"
	"github.com/OkciD/whos_on_call/internal/shared/errors"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type config struct {
	Logger    logger.Config    `json:"logger"`
	ApiClient apiclient.Config `json:"apiClient"`
	Device    *ConfigDevice    `json:"device"`
}

type deviceType string

const (
	deviceTypeLaptop deviceType = "laptop"
	deviceTypeMobile deviceType = "mobile"
	deviceTypePC     deviceType = "pc"
)

type ConfigDevice struct {
	Type deviceType `json:"type"`
	Name string     `json:"name"`
}

func (d *ConfigDevice) ToAppModel() (*models.Device, error) {
	appDevice := &models.Device{
		Name: d.Name,
	}

	switch d.Type {
	case deviceTypeLaptop:
		appDevice.Type = models.DeviceTypeLaptop
	case deviceTypeMobile:
		appDevice.Type = models.DeviceTypeMobile
	case deviceTypePC:
		appDevice.Type = models.DeviceTypePC
	default:
		return nil, errors.ErrDeviceTypeInvalid
	}

	return appDevice, nil
}
