package device

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/models"
)

type Repository interface {
	Create(ctx context.Context, newDevice *models.Device) (*models.Device, error)
}
