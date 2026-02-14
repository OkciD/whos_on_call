package httpclient

import "github.com/OkciD/whos_on_call/internal/pkg/duration"

type Config struct {
	Timeout duration.MarshallableDuration `json:"timeout"`
	ApiKey  string                        `json:"apiKey"`
}
