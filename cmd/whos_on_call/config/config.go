package config

import (
	"github.com/OkciD/whos_on_call/internal/pkg/duration"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Config struct {
	Server struct {
		ListenAddr string `json:"listenAddr"`
	} `json:"server"`

	Logger logger.Config `json:"logger"`

	DB struct {
		DSN             string                        `json:"dsn"`
		MaxIdleConns    int                           `json:"maxIdleConns"`
		MaxOpenConns    int                           `json:"maxOpenConns"`
		ConnMaxLifetime duration.MarshallableDuration `json:"connMaxLifetime"`
		ConnMaxIdleTime duration.MarshallableDuration `json:"connMaxIdleTime"`
		PingTimeout     duration.MarshallableDuration `json:"pingTimeout"`
	} `json:"db"`
}
