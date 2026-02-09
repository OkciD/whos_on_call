package api

import (
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/models"
	"github.com/OkciD/whos_on_call/internal/server/pkg/errors"
)

type callState string

const (
	callStateInactive callState = "inactive"
	callStateActive   callState = "active"
)

type deviceStatus struct {
	Device
	Features []DeviceFeature `json:"features"`
}

type userStatus struct {
	User    User           `json:"user"`
	State   callState      `json:"state"`
	Devices []deviceStatus `json:"devices"`
}

type CallStatus []userStatus

func FromAppCallStatus(appCallStatus appModels.CallStatus) (CallStatus, error) {
	apiCallStatus := make(CallStatus, 0, len(appCallStatus))

	for _, appUserStatus := range appCallStatus {
		apiUserStatus := userStatus{
			User:    *FromUserAppModel(appUserStatus.User),
			Devices: make([]deviceStatus, 0, len(appUserStatus.Devices)),
		}

		switch appUserStatus.State {
		case appModels.CallStateInactive:
			apiUserStatus.State = callStateInactive
		case appModels.CallStateActive:
			apiUserStatus.State = callStateActive
		default:
			return nil, errors.ErrCallStatusStateInvalid
		}

		for _, appDeviceStatus := range appUserStatus.Devices {
			apiDevice, err := FromDeviceAppModel(&appDeviceStatus.Device)
			if err != nil {
				return nil, fmt.Errorf("error transforming api device to app: %w", err)
			}

			apiDeviceStatus := deviceStatus{
				Device:   *apiDevice,
				Features: make([]DeviceFeature, 0, len(appDeviceStatus.Features)),
			}

			for _, appDeviceFeature := range appDeviceStatus.Features {
				apiFeature, err := FromDeviceFeatureAppModel(&appDeviceFeature)
				if err != nil {
					return nil, fmt.Errorf("failed to convert app device feature to api: %w", err)
				}

				apiDeviceStatus.Features = append(apiDeviceStatus.Features, *apiFeature)
			}

			apiUserStatus.Devices = append(apiUserStatus.Devices, apiDeviceStatus)
		}

		apiCallStatus = append(apiCallStatus, apiUserStatus)
	}

	return apiCallStatus, nil
}
