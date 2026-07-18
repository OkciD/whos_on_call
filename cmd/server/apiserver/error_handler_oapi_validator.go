package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	appErrors "github.com/OkciD/whos_on_call/internal/shared/errors"
	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func respondError(w http.ResponseWriter, errorResp mapper.ErrorResp) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(errorResp.StatusCode)

	err := json.NewEncoder(w).Encode(errorResp.ErrorResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"code\":\"internal\"}"))
		return
	}

}

// https://pkg.go.dev/github.com/oapi-codegen/nethttp-middleware#example-OapiRequestValidatorWithOptions-WithErrorHandlerWithOpts
// todo: refactor
func NewOapiValidatorErrorHandler(logger loggerPkg.Logger) nethttpmiddleware.ErrorHandlerWithOpts {
	return func(ctx context.Context, err error, w http.ResponseWriter, r *http.Request, opts nethttpmiddleware.ErrorHandlerOpts) {
		logger := logger.WithContext(ctx).WithError(err)

		if opts.MatchedRoute == nil {
			logger.Error("no matched route found")

			respondError(w, mapper.ErrorToResp(appErrors.ErrRouteRouteNotFound))
			return
		}

		switch e := err.(type) {
		case *openapi3filter.SecurityRequirementsError:
			logger.Error("security requirements error")

			respondError(w, mapper.ErrorToResp(appErrors.ErrUnauthorized))
			return
		case *openapi3filter.RequestError:
			logger.Error("request error")

			if e.RequestBody != nil && e.RequestBody.Required && r.ContentLength == 0 {
				resp := mapper.ErrorToResp(appErrors.ErrInvalid)
				resp.Body = &api.ErrorResponse_Body{}
				resp.Body.FromErrorResponseWholeRequestError(api.ErrorResponseWholeRequestErrorRequired)
				respondError(w, resp)
				return
			}

			if childErr := e.Unwrap(); childErr != nil {
				switch ce := childErr.(type) {
				case *openapi3.SchemaError:
					resp := mapper.ErrorToResp(appErrors.ErrInvalid)
					path := strings.Join(ce.JSONPointer(), ".")

					resp.Body = &api.ErrorResponse_Body{}
					if path != "" {
						fieldError := "invalid"
						if ce.SchemaField == "required" {
							fieldError = "required"
						}
						resp.Body.FromErrorResponseRequestFieldError(map[string]string{
							path: fieldError,
						})
					} else {
						resp.Body.FromErrorResponseWholeRequestError(api.ErrorResponseWholeRequestErrorInvalid)
					}

					respondError(w, resp)
					return
				case *openapi3filter.ParseError:
					p := e.Parameter

					resp := mapper.ErrorToResp(appErrors.ErrInvalid)
					if p.In == "path" {
						resp.UrlParams = &api.ErrorResponse_UrlParams{}
						resp.UrlParams.FromErrorResponseRequestFieldError(api.ErrorResponseRequestFieldError{
							p.Name: "invalid",
						})
					}
					// todo: query

					respondError(w, resp)
					return
				default:
					respondError(w, mapper.ErrorToResp(ce))
					return
				}
			}
		}

		respondError(w, mapper.ErrorToResp(appErrors.ErrInternal))
	}
}
