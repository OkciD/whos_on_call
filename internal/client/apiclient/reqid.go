package apiclient

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/OkciD/whos_on_call/internal/client/apiclient/gen"
)

const ReqIDHeader string = "X-Request-ID"

func newRequestIDEditor() gen.RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		if req == nil {
			return nil
		}

		req.Header.Add(ReqIDHeader, uuid.NewString())

		return nil
	}
}
