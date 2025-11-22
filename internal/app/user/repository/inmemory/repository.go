package inmemory

import (
	"fmt"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/google/uuid"
)

type Config struct {
	Users []struct {
		UID    string `mapstructure:"uid"`
		Name   string `mapstructure:"name"`
		ApiKey string `mapstructure:"apiKey"`
	}
}

type Repository struct {
	storage map[uuid.UUID]*models.User
}

func New(config *Config) (user.Repository, error) {
	storage := make(map[uuid.UUID]*models.User, len(config.Users))
	for _, u := range config.Users {
		uid, err := uuid.Parse(u.UID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse user uid %s: %w", uid, err)
		}
		storage[uid] = &models.User{
			UID:    uid,
			Name:   u.Name,
			ApiKey: u.ApiKey,
		}
	}

	return &Repository{
		storage: storage,
	}, nil
}
