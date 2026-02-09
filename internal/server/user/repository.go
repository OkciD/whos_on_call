package user

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/models"
)

type Repository interface {
	List(ctx context.Context) ([]models.User, error)
	GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error)
}
