package web

import "github.com/OkciD/whos_on_call/internal/shared/pkg/duration"

type Config struct {
	CallStatusPollingInterval duration.MarshallableDuration `json:"callStatusPollingInterval"`
}
