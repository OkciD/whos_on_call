package config

import (
	UserRepositoryInmemory "github.com/OkciD/whos_on_call/internal/app/user/repository/inmemory"
)

type Config struct {
	Server struct {
		ListenAddr string `json:"listenAddr"`
	} `json:"server"`

	User struct {
		Repository UserRepositoryInmemory.Config `json:"repository"`
	} `json:"user"`
}
