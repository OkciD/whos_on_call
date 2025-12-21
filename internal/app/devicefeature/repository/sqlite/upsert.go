package sqlite

import (
	"context"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/models/db"
)

func (r *Repository) Upsert(
	ctx context.Context,
	newDeviceFeature *appModels.DeviceFeature,
) (*appModels.DeviceFeature, error) {
	dbNewDeviceFeature, err := db.FromDeviceFeatureAppModel(newDeviceFeature)
	if err != nil {
		return nil, fmt.Errorf("failed device feature convertion from app to db: %w", err)
	}

	result, err := r.db.ExecContext(
		ctx,
		"INSERT OR REPLACE INTO device_features (type, status, last_modified, device_id) VALUES (?, ?, ?, ?)",
		dbNewDeviceFeature.Type,
		dbNewDeviceFeature.Status,
		dbNewDeviceFeature.LastModified,
		dbNewDeviceFeature.DeviceID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert/update device_feature: %w", err)
	}

	if lastInsertId, err := result.LastInsertId(); err == nil {
		newDeviceFeature.ID = int(lastInsertId)
	} else {
		return nil, fmt.Errorf("error getting last inserted device feature id: %w", err)
	}

	return newDeviceFeature, nil
}
