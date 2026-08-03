package apiclient

import "github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"

type Config struct {
	HTTPClientConfig httpclient.Config `json:"httpClient"`
	ApiKey           string            `json:"apiKey"`
	BaseURL          string            `json:"baseUrl"`
}
