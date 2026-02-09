package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/models"
	dbModels "github.com/OkciD/whos_on_call/internal/models/db"
	"github.com/OkciD/whos_on_call/internal/server/pkg/errors"
)

func (r *Repository) GetById(ctx context.Context, deviceID int, userID int) (*appModels.Device, error) {
	result := dbModels.Device{}

	row := r.db.QueryRowContext(ctx, "SELECT id, name, type FROM devices WHERE id = ? AND user_id = ?", deviceID, userID)

	if err := row.Scan(&result.ID, &result.Name, &result.Type); err != nil {
		r.logger.WithError(err).Error("error selecting device by id and user")

		if err == sql.ErrNoRows {
			return nil, errors.ErrDeviceNotFound
		}

		return nil, fmt.Errorf("error selecting device: %w", err)
	}

	appDevice, err := result.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("error transforming device to app model: %w", err)
	}

	return appDevice, nil
}
