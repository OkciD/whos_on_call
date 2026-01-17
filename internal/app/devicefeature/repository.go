package devicefeature

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type Repository interface {
	Create(ctx context.Context, newDeviceFeature *models.DeviceFeature) (*models.DeviceFeature, error)
	GetByDeviceId(ctx context.Context, deviceID int, deviceFeatureType models.DeviceFeatureType) (*models.DeviceFeature, error)
	Update(ctx context.Context, updatedDeviceFeature *models.DeviceFeature) error
	ListByDeviceID(ctx context.Context, deviceID int) ([]models.DeviceFeature, error)
}
