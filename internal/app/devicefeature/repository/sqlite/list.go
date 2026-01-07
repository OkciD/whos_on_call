package sqlite

import (
	"context"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	dbModels "github.com/OkciD/whos_on_call/internal/app/models/db"
)

func (r *Repository) ListByDeviceID(ctx context.Context, deviceID int) ([]appModels.DeviceFeature, error) {
	dbFeatures := make([]dbModels.DeviceFeature, 0, 2)

	rows, err := r.db.QueryContext(ctx, "SELECT id, type, status, last_modified FROM device_features WHERE device_id = ?", deviceID)
	if err != nil {
		return nil, fmt.Errorf("select device features by device query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		dbFeature := dbModels.DeviceFeature{}
		if err := rows.Scan(&dbFeature.ID, &dbFeature.Type, &dbFeature.Status, &dbFeature.LastModified); err != nil {
			return nil, fmt.Errorf("failed to scan db feature from row: %w", err)
		}
		dbFeatures = append(dbFeatures, dbFeature)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in features by device id query: %w", err)
	}

	appFeatures := make([]appModels.DeviceFeature, 0, len(dbFeatures))
	for _, dbFeature := range dbFeatures {
		appFeature, err := dbFeature.ToAppModel()
		if err != nil {
			return nil, fmt.Errorf("error converting device feature from db to app: %w", err)
		}

		appFeatures = append(appFeatures, *appFeature)
	}

	return appFeatures, nil
}
