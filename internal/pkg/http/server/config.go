package server

import "github.com/OkciD/whos_on_call/internal/pkg/duration"

type Config struct {
	ListenAddr      string                        `json:"listenAddr"`
	ShutdownTimeout duration.MarshallableDuration `json:"shutdownTimeout"`
}
