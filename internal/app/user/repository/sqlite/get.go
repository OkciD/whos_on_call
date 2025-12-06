package inmemory

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	appModels "github.com/OkciD/whos_on_call/internal/app/models"
	dbModels "github.com/OkciD/whos_on_call/internal/app/models/db"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
)

func (r *Repository) GetUserByApiKey(ctx context.Context, apiKey string) (*appModels.User, error) {
	hasher := sha256.New()
	hasher.Write([]byte(apiKey))
	hashedApiKey := hex.EncodeToString(hasher.Sum(nil))

	r.logger.WithField("api_key_hash", hashedApiKey).Trace("lookup user by api key")

	user := dbModels.User{}

	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE api_key_hash = ?", hashedApiKey)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		r.logger.WithError(err).Error("error selecting user")

		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}

		return nil, fmt.Errorf("error selecting user: %w", err)
	}

	return user.ToAppModel(), nil
}
