package apiclient

import (
	"context"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/client/apiclient/gen"
)

//nolint:gosec // фолзит
const APIKeyHeader = "X-Api-Key"

func newAuthRequestEditor(apiKey string) gen.RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		if req == nil {
			return nil
		}

		req.Header.Add(APIKeyHeader, apiKey)

		return nil
	}
}
