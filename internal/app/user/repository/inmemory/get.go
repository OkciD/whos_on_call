package inmemory

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
)

func (r *Repository) GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error) {
	for _, user := range r.storage {
		if user.ApiKey == apiKey {
			return user, nil
		}
	}

	return nil, errors.ErrUserNotFound
}
