package device

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type UseCase interface {
	Create(ctx context.Context, newDevice *models.Device) (*models.Device, error)
}
