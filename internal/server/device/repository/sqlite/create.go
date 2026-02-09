package sqlite

import (
	"context"
	"errors"

	"fmt"

	"github.com/OkciD/whos_on_call/internal/models"
	"github.com/OkciD/whos_on_call/internal/models/db"
	appErrors "github.com/OkciD/whos_on_call/internal/server/pkg/errors"
	sqlite "github.com/mattn/go-sqlite3"
)

func (r *Repository) Create(ctx context.Context, newDevice *models.Device) (*models.Device, error) {
	dbNewDevice, err := db.FromDeviceAppModel(newDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device model for creation: %w", err)
	}

	result, err := r.db.ExecContext(ctx, "INSERT INTO devices (name, type, user_id) VALUES (?, ?, ?)", dbNewDevice.Name, dbNewDevice.Type, dbNewDevice.UserID)
	if err != nil {
		if sqliteError, ok := err.(sqlite.Error); ok {
			if sqliteError.Code == sqlite.ErrConstraint && sqliteError.ExtendedCode == sqlite.ErrConstraintUnique {
				return nil, fmt.Errorf("%w, %w", appErrors.ErrDeviceExists, errors.New(sqliteError.Error()))
			}
		}

		return nil, fmt.Errorf("failed to insert device: %w", err)
	}

	if deviceID, err := result.LastInsertId(); err == nil {
		newDevice.ID = int(deviceID)
	} else {
		return nil, fmt.Errorf("failed to get inserted device id: %w", err)
	}

	return newDevice, nil
}
