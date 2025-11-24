package config

import (
	UserRepositoryInmemory "github.com/OkciD/whos_on_call/internal/app/user/repository/inmemory"
)

type LogFormat string

const (
	LogFormatJson LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type Config struct {
	Server struct {
		ListenAddr string `json:"listenAddr"`
	} `json:"server"`

	Logger struct {
		Level  string    `json:"level"`
		Format LogFormat `json:"format"`
	} `json:"logger"`

	User struct {
		Repository UserRepositoryInmemory.Config `json:"repository"`
	} `json:"user"`
}

func (c *Config) GetLogFormat() LogFormat {
	if c.Logger.Format == "" {
		return LogFormatJson
	}

	return c.Logger.Format
}
