package sqlite

import (
	"context"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	dbModels "github.com/OkciD/whos_on_call/internal/app/models/db"
)

func (r *Repository) ListByUserID(ctx context.Context, userID int) ([]appModels.Device, error) {
	dbDevices := make([]dbModels.Device, 0, 4)

	rows, err := r.db.QueryContext(ctx, "SELECT id, name, type FROM devices WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("select devices by user query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		dbDevice := dbModels.Device{}
		if err := rows.Scan(&dbDevice.ID, &dbDevice.Name, &dbDevice.Type); err != nil {
			return nil, fmt.Errorf("failed to scan db device from row: %w", err)
		}
		dbDevices = append(dbDevices, dbDevice)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in devices by user id query: %w", err)
	}

	appDevices := make([]appModels.Device, 0, len(dbDevices))
	for _, dbDevice := range dbDevices {
		appDevice, err := dbDevice.ToAppModel()
		if err != nil {
			return nil, fmt.Errorf("error converting device from db to app: %w", err)
		}

		appDevices = append(appDevices, *appDevice)
	}

	return appDevices, nil
}
