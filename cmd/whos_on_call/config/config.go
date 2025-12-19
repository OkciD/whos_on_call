package config

import (
	"github.com/OkciD/whos_on_call/internal/pkg/db"
	"github.com/OkciD/whos_on_call/internal/pkg/duration"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Config struct {
	Server struct {
		ListenAddr      string                        `json:"listenAddr"`
		ShutdownTimeout duration.MarshallableDuration `json:"shutdownTimeout"`
	} `json:"server"`

	Logger logger.Config `json:"logger"`

	DB db.Config `json:"db"`
}
