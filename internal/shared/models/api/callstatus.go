package api

import (
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/errors"
	appModels "github.com/OkciD/whos_on_call/internal/shared/models"
)

func FromAppCallStatus(appCallStatus appModels.CallStatus) (CallStatus, error) {
	apiCallStatus := make(CallStatus, 0, len(appCallStatus))

	for _, appUserStatus := range appCallStatus {
		apiUserStatus := UserStatus{
			User:    *FromUserAppModel(appUserStatus.User),
			Devices: make([]DeviceStatus, 0, len(appUserStatus.Devices)),
		}

		switch appUserStatus.State {
		case appModels.CallStateInactive:
			apiUserStatus.State = CallStateInactive
		case appModels.CallStateActive:
			apiUserStatus.State = CallStateActive
		default:
			return nil, errors.ErrCallStatusStateInvalid
		}

		for _, appDeviceStatus := range appUserStatus.Devices {
			apiDevice, err := FromDeviceAppModel(&appDeviceStatus.Device)
			if err != nil {
				return nil, fmt.Errorf("error transforming api device to app: %w", err)
			}

			apiDeviceStatus := DeviceStatus{
				Id:       int32(apiDevice.Id),
				Name:     apiDevice.Name,
				Type:     DeviceType(apiDevice.Type),
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
