package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	appErrors "github.com/OkciD/whos_on_call/internal/errors"
	appModels "github.com/OkciD/whos_on_call/internal/models"
	dbModels "github.com/OkciD/whos_on_call/internal/models/db"
)

func (r *Repository) GetById(ctx context.Context, deviceID int, userID int) (*appModels.Device, error) {
	result := dbModels.Device{}

	row := r.db.QueryRowContext(ctx, "SELECT id, name, type FROM devices WHERE id = ? AND user_id = ?", deviceID, userID)

	if err := row.Scan(&result.ID, &result.Name, &result.Type); err != nil {
		r.logger.WithError(err).Error("error selecting device by id and user")

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: no sql rows for device %d", appErrors.ErrNotFound, deviceID)
		}

		return nil, fmt.Errorf("error selecting device: %w", err)
	}

	appDevice, err := result.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("error transforming device to app model: %w", err)
	}

	return appDevice, nil
}
