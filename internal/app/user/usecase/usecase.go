package usecase

import "github.com/OkciD/whos_on_call/internal/app/user"

type UseCase struct {
	userRepo user.Repository
}

func New(userRepo user.Repository) user.UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}
