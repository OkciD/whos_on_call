package sqlite

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/errors"
	appModels "github.com/OkciD/whos_on_call/internal/models"
	dbModels "github.com/OkciD/whos_on_call/internal/models/db"
)

func (r *Repository) GetUserByApiKey(ctx context.Context, apiKey string) (*appModels.User, error) {
	hasher := sha256.New()
	hasher.Write([]byte(apiKey))
	hashedApiKey := hex.EncodeToString(hasher.Sum(nil))

	user := dbModels.User{}

	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE api_key_hash = ?", hashedApiKey)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		r.logger.WithError(err).Error("error selecting user")

		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}

		return nil, fmt.Errorf("error selecting user: %w", err)
	}

	return user.ToAppModel(), nil
}
