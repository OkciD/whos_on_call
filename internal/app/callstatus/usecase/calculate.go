package usecase

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

func (u *UseCase) Calculate(ctx context.Context) (models.CallStatus, error) {
	users, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users to define status: %w", err)
	}

	status := make(models.CallStatus, 0, len(users))
	for _, user := range users {
		devices, err := u.deviceRepo.ListByUserID(ctx, user.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to list devices by user id %d: %w", user.ID, err)
		}

		userStatus := models.UserStatus{
			User:    &user,
			State:   models.CallStateInactive,
			Devices: make([]models.DeviceStatus, 0, len(devices)),
		}

		for _, device := range devices {
			features, err := u.deviceFeatureRepo.ListByDeviceID(ctx, device.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to list features by device id %d: %w", device.ID, err)
			}

			deviceStatus := models.DeviceStatus{
				Device:   device,
				Features: features,
			}

			if userStatus.State == models.CallStateInactive {
				for _, f := range deviceStatus.Features {
					if f.Status == models.DeviceFeatureStatusActive || f.WasActiveRecently(u.config.RelaxationPeriod.Duration) {
						userStatus.State = models.CallStateActive
						break
					}
				}
			}

			userStatus.Devices = append(userStatus.Devices, deviceStatus)
		}

		status = append(status, userStatus)
	}

	return status, nil
}
