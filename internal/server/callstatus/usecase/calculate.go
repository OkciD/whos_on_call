package usecase

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

func (u *UseCase) Calculate(ctx context.Context) (models.CallStatus, error) {
	users, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users to define status: %w", err)
	}

	status := make(models.CallStatus, len(users))
	g, gctx := errgroup.WithContext(ctx)

	for i, user := range users {
		g.Go(func() error {
			userStatus, err := u.getUserStatus(gctx, &user)
			if err != nil {
				return fmt.Errorf("error getting user (id: %d) status : %w", user.ID, err)
			}
			status[i] = *userStatus
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("error getting call status: %w", err)
	}

	return status, nil
}

func (u *UseCase) getUserStatus(ctx context.Context, user *models.User) (*models.UserStatus, error) {
	devices, err := u.deviceRepo.ListByUserID(ctx, user.ID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list devices by user id %d: %w", user.ID, err)
	}

	userStatus := models.UserStatus{
		User:    user,
		State:   models.CallStateInactive,
		Devices: make([]models.DeviceStatus, len(devices)),
	}

	g, gctx := errgroup.WithContext(ctx)

	for i, device := range devices {
		g.Go(func() error {
			deviceStatus, err := u.getDeviceStatus(gctx, &device)
			if err != nil {
				return fmt.Errorf("error getting device (id: %d) status: %w", device.ID, err)
			}
			userStatus.Devices[i] = *deviceStatus
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("error getting user status: %w", err)
	}

	for _, deviceStatus := range userStatus.Devices {
		for _, f := range deviceStatus.Features {
			if f.Status == models.DeviceFeatureStatusActive ||
				f.WasActiveRecently(u.config.RelaxationPeriod.Duration) {
				userStatus.State = models.CallStateActive
				break
			}
		}
	}

	return &userStatus, nil
}

func (u *UseCase) getDeviceStatus(ctx context.Context, device *models.Device) (*models.DeviceStatus, error) {
	features, err := u.deviceFeatureRepo.ListByDeviceID(ctx, device.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to list features by device id %d: %w", device.ID, err)
	}

	return &models.DeviceStatus{
		Device:   *device,
		Features: features,
	}, nil
}
