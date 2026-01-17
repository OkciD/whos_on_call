package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OkciD/whos_on_call/internal/app/models"
	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func (u *UseCase) Upsert(
	ctx context.Context,
	deviceID int,
	user *models.User,
	newDeviceFeature *models.DeviceFeature,
) (*models.DeviceFeature, error) {
	u.logger.WithFields(logger.Fields{
		"deviceID":               deviceID,
		"user":                   user,
		"newDeviceFeatureType":   newDeviceFeature.Type,
		"newDeviceFeatureStatus": newDeviceFeature.Status,
	}).Info("upsert device feature")

	device, err := u.deviceRepo.GetById(ctx, deviceID, user.ID)
	device.User = user
	if err != nil {
		return nil, fmt.Errorf("failed to get device from repo: %w", err)
	}
	u.logger.WithFields(logger.Fields{
		"deviceID":   device.ID,
		"deviceName": device.Name,
		"deviceType": device.Type,
	}).Info("device found")

	existingDeviceFeature, err := u.deviceFeatureRepo.GetByDeviceId(ctx, device.ID, newDeviceFeature.Type)
	if err != nil && !errors.Is(err, appErrors.ErrDeviceFeatureNotFound) {
		return nil, fmt.Errorf("error checking existing device feature in db: %w", err)
	}

	if existingDeviceFeature != nil {
		u.logger.WithFields(logger.Fields{
			"deviceID":            deviceID,
			"deviceFeatureID":     existingDeviceFeature.ID,
			"deviceFeatureType":   newDeviceFeature.Type,
			"deviceFeatureStatus": newDeviceFeature.Status,
		}).Info("device feature already exists in db")
	} else {
		u.logger.WithFields(logger.Fields{
			"deviceID":          deviceID,
			"deviceFeatureType": newDeviceFeature.Type,
		}).Info("no existing device feature found")
	}

	newDeviceFeature.Device = device

	if newDeviceFeature.Status == models.DeviceFeatureStatusActive {
		tn := time.Now()
		newDeviceFeature.LastActive = &tn
	} else if existingDeviceFeature != nil {
		newDeviceFeature.LastActive = existingDeviceFeature.LastActive
	}

	if existingDeviceFeature != nil {
		newDeviceFeature.ID = existingDeviceFeature.ID

		err = u.deviceFeatureRepo.Update(ctx, newDeviceFeature)
		if err != nil {
			return nil, fmt.Errorf("failed to update device feature in repo: %w", err)
		}

		u.logger.WithFields(logger.Fields{
			"deviceFeatureID":     newDeviceFeature.ID,
			"deviceFeatureType":   newDeviceFeature.Type,
			"deviceFeatureStatus": newDeviceFeature.Status,
		}).Info("device feature updated successfully")

		return newDeviceFeature, nil
	} else {
		newDeviceFeature, err := u.deviceFeatureRepo.Create(ctx, newDeviceFeature)
		if err != nil {
			return nil, fmt.Errorf("failed to create device feature in repo: %w", err)
		}

		u.logger.WithFields(logger.Fields{
			"deviceFeatureID":     newDeviceFeature.ID,
			"deviceFeatureType":   newDeviceFeature.Type,
			"deviceFeatureStatus": newDeviceFeature.Status,
		}).Info("device feature created successfully")

		return newDeviceFeature, nil
	}
}
