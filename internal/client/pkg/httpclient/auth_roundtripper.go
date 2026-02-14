package httpclient

import (
	"net/http"
)

type authRoundTripper struct {
	next   http.RoundTripper
	apiKey string
}

const API_KEY_HEADER = "X-Api-Key"

func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add(API_KEY_HEADER, a.apiKey)

	resp, err := a.next.RoundTrip(r)

	return resp, err
}
