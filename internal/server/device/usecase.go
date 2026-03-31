package device

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type UseCase interface {
	Create(ctx context.Context, newDevice *models.Device) (*models.Device, error)
}
