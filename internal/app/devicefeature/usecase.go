package devicefeature

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type UseCase interface {
	Upsert(
		ctx context.Context,
		deviceId int,
		user *models.User,
		newDeviceFeature *models.DeviceFeature,
	) (*models.DeviceFeature, error)
}
