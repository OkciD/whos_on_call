package usecase

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (u *UseCase) List(ctx context.Context, user *models.User, params *models.DeviceListParams) ([]models.Device, error) {
	u.logger.WithContext(ctx).WithField("params", params).Info("list devices")

	devices, err := u.deviceRepo.ListByUserID(ctx, user.ID, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices list from repo: %w", err)
	}

	return devices, nil
}
