package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func (u *UseCase) Upsert(
	ctx context.Context,
	deviceId int,
	user *models.User,
	newDeviceFeature *models.DeviceFeature,
) (*models.DeviceFeature, error) {
	u.logger.WithFields(logger.Fields{
		"deviceId":       deviceId,
		"user":           user,
		"device_feature": newDeviceFeature,
	}).Info("upsert device feature")

	device, err := u.deviceRepo.GetById(ctx, deviceId, user)
	if err != nil {
		return nil, fmt.Errorf("failed to get device from repo: %w", err)
	}

	newDeviceFeature.Device = device

	tn := time.Now()
	newDeviceFeature.LastModified = &tn

	newDeviceFeature, err = u.deviceFeatureRepo.Upsert(ctx, newDeviceFeature)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert device feature in repo: %w", err)
	}

	return newDeviceFeature, nil
}
