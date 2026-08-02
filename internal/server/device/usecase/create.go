package usecase

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (u *UseCase) Create(ctx context.Context, newDevice *models.Device) (*models.Device, error) {
	err := u.txManager.WithinTransaction(ctx, func(txctx context.Context) error {
		createdDevice, err := u.deviceRepo.Create(txctx, newDevice)
		if err != nil {
			return fmt.Errorf("failed to create device in repo: %w", err)
		}

		for _, featureType := range models.DeviceFeatureTypes {
			// todo: log
			_, err := u.deviceFeatureRepo.Create(txctx, &models.DeviceFeature{
				Type:   featureType,
				Status: models.DeviceFeatureStatusInactive,
				Device: createdDevice,
			})
			if err != nil {
				return fmt.Errorf("failed to create feature for device %d in repo: %w", createdDevice.ID, err)
			}
		}

		newDevice = createdDevice

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create device and features in repo: %w", err)
	}

	return newDevice, nil
}
