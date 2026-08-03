package httpclient

import "github.com/OkciD/whos_on_call/internal/shared/pkg/duration"

type Config struct {
	Timeout duration.MarshallableDuration `json:"timeout"`

	Logging struct {
		Request struct {
			HideHeaders []string `json:"hideHeaders"`
		} `json:"request"`
	} `json:"logging"`
}
