package usecase

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/models"
)

func (u *UseCase) Create(ctx context.Context, newDevice *models.Device) (*models.Device, error) {
	newDevice, err := u.deviceRepo.Create(ctx, newDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to create device in repo: %w", err)
	}

	return newDevice, err
}
