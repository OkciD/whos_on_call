package config

import (
	"github.com/OkciD/whos_on_call/internal/pkg/duration"
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

	DB struct {
		DSN             string                        `json:"dsn"`
		MaxIdleConns    int                           `json:"maxIdleConns"`
		MaxOpenConns    int                           `json:"maxOpenConns"`
		ConnMaxLifetime duration.MarshallableDuration `json:"connMaxLifetime"`
		ConnMaxIdleTime duration.MarshallableDuration `json:"connMaxIdleTime"`
		PingTimeout     duration.MarshallableDuration `json:"pingTimeout"`
	} `json:"db"`
}

func (c *Config) GetLogFormat() LogFormat {
	if c.Logger.Format == "" {
		return LogFormatJson
	}

	return c.Logger.Format
}
