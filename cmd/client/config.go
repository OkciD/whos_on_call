package main

import (
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type config struct {
	Logger logger.Config `json:"logger"`
}
