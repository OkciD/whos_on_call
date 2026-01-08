package config

import (
	callStatusUseCase "github.com/OkciD/whos_on_call/internal/app/callstatus/usecase"
	"github.com/OkciD/whos_on_call/internal/pkg/db"
	"github.com/OkciD/whos_on_call/internal/pkg/http/server"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Config struct {
	ApiServer server.Config `json:"apiServer"`
	WebServer server.Config `json:"webServer"`

	Logger logger.Config `json:"logger"`

	CallStatus struct {
		UseCase callStatusUseCase.Config `json:"useCase"`
	} `json:"callStatus"`

	DB db.Config `json:"db"`
}
