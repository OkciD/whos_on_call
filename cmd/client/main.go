package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/OkciD/whos_on_call/internal/client/apiclient"
	appErrors "github.com/OkciD/whos_on_call/internal/shared/errors"
	configUtils "github.com/OkciD/whos_on_call/internal/shared/pkg/config"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func main() {
	configFilePathPtr := flag.String("config", "", "path to config file")

	flag.Parse()

	if configFilePathPtr == nil || *configFilePathPtr == "" {
		log.Fatal("-config option required")
	}

	cfg, err := configUtils.ReadConfig[config](*configFilePathPtr)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config: %w", err))
	}

	logger := loggerPkg.NewLogrusBasedLogger(&cfg.Logger)

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
