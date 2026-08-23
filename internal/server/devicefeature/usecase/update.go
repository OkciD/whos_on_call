package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func (u *UseCase) Update(
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

	device, err := u.deviceRepo.GetByID(ctx, deviceID, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get device from repo: %w", err)
	}
	device.User = user

	u.logger.WithFields(logger.Fields{
		"deviceID":   device.ID,
		"deviceName": device.Name,
		"deviceType": device.Type,
	}).Info("device found")

	existingDeviceFeature, err := u.deviceFeatureRepo.GetByDeviceID(ctx, device.ID, newDeviceFeature.Type)
	if err != nil {
		return nil, fmt.Errorf("error checking existing device feature in db: %w", err)
	}
	u.logger.WithFields(logger.Fields{
		"deviceID":            deviceID,
		"deviceFeatureID":     existingDeviceFeature.ID,
		"deviceFeatureType":   existingDeviceFeature.Type,
		"deviceFeatureStatus": existingDeviceFeature.Status,
	}).Info("device feature found in db")

	newDeviceFeature.Device = device

	if newDeviceFeature.Status == models.DeviceFeatureStatusActive {
		newDeviceFeature.LastActive = new(time.Now())
	} else if existingDeviceFeature != nil {
		newDeviceFeature.LastActive = existingDeviceFeature.LastActive
	}

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
}
