package usecase

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/models"
)

func (u *UseCase) GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error) {
	u.logger.WithContext(ctx).Info("get user by api key")

	return u.userRepo.GetUserByApiKey(ctx, apiKey)
}
