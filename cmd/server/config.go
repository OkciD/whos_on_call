package main

import (
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	callStatusUseCase "github.com/OkciD/whos_on_call/internal/server/callstatus/usecase"
	"github.com/OkciD/whos_on_call/internal/server/pkg/db"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/server"
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
