package devicefeature

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type Repository interface {
	Upsert(ctx context.Context, newDeviceFeature *models.DeviceFeature) (*models.DeviceFeature, error)
	ListByDeviceID(ctx context.Context, deviceID int) ([]models.DeviceFeature, error)
}
