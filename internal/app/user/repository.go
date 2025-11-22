package user

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type Repository interface {
	GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error)
}
