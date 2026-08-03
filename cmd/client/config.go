package main

import (
	"github.com/OkciD/whos_on_call/internal/client/pkg/apiclient"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type config struct {
	Logger    logger.Config    `json:"logger"`
	ApiClient apiclient.Config `json:"apiClient"`
}
