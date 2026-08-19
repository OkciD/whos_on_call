package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type params struct {
	configPath          string
	deviceFeatureType   models.DeviceFeatureType
	deviceFeatureStatus models.DeviceFeatureStatus
}

// OverSight-compatible flags
// https://github.com/objective-see/OverSight
func readFlags() (*params, error) {
	fs := flag.NewFlagSet("whos_on_call", flag.ContinueOnError)

	configFilePathPtr := fs.String("config", "", "path to config file")
	deviceFeatureTypePtr := fs.String("device", "", "device feature (mic/camera)")
	eventPtr := fs.String("event", "", "device event (on/off)")

	fs.Parse(os.Args[1:])

	params := params{}

	if *configFilePathPtr == "" {
		currentPath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("failed to get executable path: %w", err)
		}

		currentDir, _ := filepath.Split(currentPath)
		params.configPath = filepath.Join(currentDir, "config.json")
	} else {
		params.configPath = *configFilePathPtr
	}

	switch *deviceFeatureTypePtr {
	case "camera":
		params.deviceFeatureType = models.DeviceFeatureTypeCamera
	case "microphone":
		params.deviceFeatureType = models.DeviceFeatureTypeMic
	case "":
		return nil, errors.New("-device param required")
	default:
		return nil, errors.New("invalid -device param")
	}

	switch *eventPtr {
	case "on":
		params.deviceFeatureStatus = models.DeviceFeatureStatusActive
	case "off":
		params.deviceFeatureStatus = models.DeviceFeatureStatusInactive
	case "":
		return nil, errors.New("-event param required")
	default:
		return nil, errors.New("invalid -event param")
	}

	return &params, nil
}
