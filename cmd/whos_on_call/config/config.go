package config

import (
	callStatusUseCase "github.com/OkciD/whos_on_call/internal/app/callstatus/usecase"
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

	CallStatus struct {
		UseCase callStatusUseCase.Config `json:"useCase"`
	} `json:"callStatus"`

	DB db.Config `json:"db"`
}
