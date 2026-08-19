package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/OkciD/whos_on_call/internal/client/apiclient"
	appErrors "github.com/OkciD/whos_on_call/internal/shared/errors"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	configUtils "github.com/OkciD/whos_on_call/internal/shared/pkg/config"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func main() {
	flags, err := readFlags()
	if err != nil {
		log.Fatalf("failed to parse flags: %s", err)
	}

	cfg, err := configUtils.ReadConfig[config](flags.configPath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config: %w", err))
	}

	logger := loggerPkg.NewLogrusBasedLogger(&cfg.Logger)

	logger.WithFields(loggerPkg.Fields{
		"configPath":          flags.configPath,
		"deviceFeatureType":   flags.deviceFeatureType,
		"deviceFeatureStatus": flags.deviceFeatureStatus,
	}).Info("flags read successfully")

	apiClient, err := apiclient.New(
		logger.ForModule("api_client"),
		cfg.ApiClient,
	)
	if err != nil {
		log.Fatal(err)
	}

	user, err := apiClient.GetUser(context.TODO())
	if err != nil {
		logger.WithError(err).Fatal("failed to get user")
	}
	logger.WithFields(loggerPkg.Fields{
		"id":   user.ID,
		"name": user.Name,
	}).Info("got user")

	appDevice, err := cfg.Device.ToAppModel()
	if err != nil {
		logger.WithError(err).Fatal("failed to get device from config")
	}

	existingDevices, err := apiClient.ListDevices(context.TODO(), &models.DeviceListParams{
		Type: &appDevice.Type,
		Name: &appDevice.Name,
	})
	if err != nil {
		logger.WithError(err).Fatal("failed to list devices")
	}

	if existingDevices != nil {
		if len(existingDevices) != 1 {
			logger.WithError(err).WithFields(loggerPkg.Fields{
				"type":       appDevice.Type,
				"name":       appDevice.Name,
				"devicesLen": len(existingDevices),
			}).Fatal("invalid number of devices returned by params")
		}

		appDevice.ID = existingDevices[0].ID

		logger.WithFields(loggerPkg.Fields{
			"id":   appDevice.ID,
			"name": appDevice.Name,
			"type": appDevice.Type,
		}).Info("device id obtained successfully")
	}

	if appDevice.ID == 0 {
		appDevice, err = apiClient.CreateDevice(context.TODO(), appDevice)
		if err != nil {
			if errors.Is(err, appErrors.ErrDuplicate) {
				logger.WithError(err).WithFields(loggerPkg.Fields{
					"name": appDevice.Name,
					"type": appDevice.Type,
				}).Error("device already exists")
			} else {
				logger.WithError(err).Fatal("failed to create device")
			}
		}

		logger.WithFields(loggerPkg.Fields{
			"id":   appDevice.ID,
			"name": appDevice.Name,
			"type": appDevice.Type,
		}).Info("device successfully created")
	}

	newFeature, err := apiClient.UpdateDeviceFeature(context.TODO(), &models.DeviceFeature{
		Type:   flags.deviceFeatureType,
		Status: flags.deviceFeatureStatus,
		Device: appDevice,
	})
	if err != nil {
		logger.WithError(err).Error("error updating device feature")
	}

	logger.WithFields(loggerPkg.Fields{
		"deviceFeatureId":     newFeature.ID,
		"deviceFeatureType":   newFeature.Type,
		"deviceFeatureStatus": newFeature.Status,
	}).Info("device feature updated successfully")
}
