package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/models"
	dbModels "github.com/OkciD/whos_on_call/internal/models/db"
	"github.com/OkciD/whos_on_call/internal/server/pkg/errors"
)

func (r *Repository) GetByDeviceId(
	ctx context.Context,
	deviceID int,
	deviceFeatureType appModels.DeviceFeatureType,
) (*appModels.DeviceFeature, error) {
	result := dbModels.DeviceFeature{}

	row := r.db.QueryRowContext(ctx, "SELECT id, type, status, last_active, device_id FROM device_features WHERE device_id = ? AND type = ?", deviceID, deviceFeatureType)

	if err := row.Scan(&result.ID, &result.Type, &result.Status, &result.LastActive, &result.DeviceID); err != nil {
		r.logger.WithError(err).Error("error selecting device feature by device id")

		if err == sql.ErrNoRows {
			return nil, errors.ErrDeviceFeatureNotFound
		}

		return nil, fmt.Errorf("error selecting device feature: %w", err)
	}

	appDeviceFeature, err := result.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("error transforming device feature to app model: %w", err)
	}

	return appDeviceFeature, nil
}
