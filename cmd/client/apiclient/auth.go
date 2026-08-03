package apiclient

import (
	"context"
	"net/http"

	"github.com/OkciD/whos_on_call/cmd/client/apiclient/gen"
)

const API_KEY_HEADER = "X-Api-Key"

func NewAuthRequestEditor(apiKey string) gen.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		if req == nil {
			return nil
		}

		req.Header.Add(API_KEY_HEADER, apiKey)

		return nil
	}
}
