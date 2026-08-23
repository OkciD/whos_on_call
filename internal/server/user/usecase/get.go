package usecase

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (u *UseCase) GetUserByAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	u.logger.WithContext(ctx).Info("get user by api key")

	user, err := u.userRepo.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from repo: %w", err)
	}

	return user, nil
}
