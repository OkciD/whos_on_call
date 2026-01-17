package sqlite

import (
	"context"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/models/db"
)

func (r *Repository) Update(
	ctx context.Context,
	updatedDeviceFeature *appModels.DeviceFeature,
) error {
	dbUpdatedDeviceFeature, err := db.FromDeviceFeatureAppModel(updatedDeviceFeature)
	if err != nil {
		return fmt.Errorf("failed device feature convertion from app to db: %w", err)
	}

	_, err = r.db.ExecContext(
		ctx,
		"UPDATE device_features SET status = ?, last_active = ? WHERE id = ?",
		dbUpdatedDeviceFeature.Status,
		dbUpdatedDeviceFeature.LastActive,
		dbUpdatedDeviceFeature.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update device_feature: %w", err)
	}

	return nil
}
