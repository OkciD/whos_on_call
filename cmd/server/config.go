package main

import (
	callStatusUseCase "github.com/OkciD/whos_on_call/internal/server/callstatus/usecase"
	"github.com/OkciD/whos_on_call/internal/server/pkg/db"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/server"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type config struct {
	ApiServer server.Config `json:"apiServer"`
	WebServer server.Config `json:"webServer"`

	Logger logger.Config `json:"logger"`

	CallStatus struct {
		UseCase callStatusUseCase.Config `json:"useCase"`
	} `json:"callStatus"`

	DB db.Config `json:"db"`
}
