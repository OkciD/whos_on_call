package server

import "github.com/OkciD/whos_on_call/internal/shared/pkg/duration"

type Config struct {
	ListenAddr        string                        `json:"listenAddr"`
	ShutdownTimeout   duration.MarshallableDuration `json:"shutdownTimeout"`
	ReadTimeout       duration.MarshallableDuration `json:"readTimeout"`
	ReadHeaderTimeout duration.MarshallableDuration `json:"readHeaderTimeout"`
	WriteTimeout      duration.MarshallableDuration `json:"writeTimeout"`
	IdleTimeout       duration.MarshallableDuration `json:"idleTimeout"`
}
