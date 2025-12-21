package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	dbModels "github.com/OkciD/whos_on_call/internal/app/models/db"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
)

func (r *Repository) GetById(ctx context.Context, deviceId int, user *appModels.User) (*appModels.Device, error) {
	result := dbModels.Device{}

	row := r.db.QueryRowContext(ctx, "SELECT id, name, type FROM devices WHERE id = ?", deviceId)

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
	appDevice.User = user

	return appDevice, nil
}
