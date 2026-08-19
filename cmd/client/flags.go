package main

import (
	"errors"
	"flag"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type deviceFeatureEvent string

const (
	deviceFeatureEventOn  deviceFeatureEvent = "on"
	deviceFeatureEventOff deviceFeatureEvent = "off"
)

type params struct {
	configPath        string
	deviceFeatureType models.DeviceFeatureType
	event             deviceFeatureEvent
}

// OverSight-compatible flags
// https://github.com/objective-see/OverSight
func readFlags() (*params, error) {
	configFilePathPtr := flag.String("config", "./config.json", "path to config file")
	deviceFeatureTypePtr := flag.String("device", "", "device feature (mic/camera)")
	eventPtr := flag.String("event", "", "device event (on/off)")

	flag.Parse()

	params := params{
		configPath: *configFilePathPtr,
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
		params.event = deviceFeatureEventOn
	case "off":
		params.event = deviceFeatureEventOff
	case "":
		return nil, errors.New("-event param required")
	default:
		return nil, errors.New("invalid -event param")
	}

	return &params, nil
}
