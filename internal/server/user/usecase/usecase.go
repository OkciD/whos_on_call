package usecase

import (
	"github.com/OkciD/whos_on_call/internal/server/user"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type UseCase struct {
	logger logger.Logger

	userRepo user.Repository
}

func New(logger logger.Logger, userRepo user.Repository) user.UseCase {
	return &UseCase{
		logger: logger,

		userRepo: userRepo,
	}
}
