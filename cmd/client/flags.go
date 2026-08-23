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

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	p := params{}

	if *configFilePathPtr == "" {
		currentPath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("failed to get executable path: %w", err)
		}

		currentDir, _ := filepath.Split(currentPath)
		p.configPath = filepath.Join(currentDir, "config.json")
	} else {
		p.configPath = *configFilePathPtr
	}

	switch *deviceFeatureTypePtr {
	case "camera":
		p.deviceFeatureType = models.DeviceFeatureTypeCamera
	case "microphone":
		p.deviceFeatureType = models.DeviceFeatureTypeMic
	case "":
		return nil, errors.New("-device param required")
	default:
		return nil, errors.New("invalid -device param")
	}

	switch *eventPtr {
	case "on":
		p.deviceFeatureStatus = models.DeviceFeatureStatusActive
	case "off":
		p.deviceFeatureStatus = models.DeviceFeatureStatusInactive
	case "":
		return nil, errors.New("-event param required")
	default:
		return nil, errors.New("invalid -event param")
	}

	return &p, nil
}
