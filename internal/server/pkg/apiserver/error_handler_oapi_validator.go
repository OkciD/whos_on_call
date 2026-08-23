package apiserver

import (
	"context"
	"encoding/json"
	"errors"
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
		_, _ = w.Write([]byte("{\"code\":\"internal\"}"))
		return
	}
}

// https://pkg.go.dev/github.com/oapi-codegen/nethttp-middleware#example-OapiRequestValidatorWithOptions-WithErrorHandlerWithOpts

//nolint:gocognit // todo: refactor
func NewOapiValidatorErrorHandler(logger loggerPkg.Logger) nethttpmiddleware.ErrorHandlerWithOpts {
	return func(ctx context.Context, err error, w http.ResponseWriter, r *http.Request, opts nethttpmiddleware.ErrorHandlerOpts) {
		logger := logger.WithContext(ctx).WithError(err)

		if opts.MatchedRoute == nil {
			logger.Error("no matched route found")

			respondError(w, mapper.ErrorToResp(appErrors.ErrRouteRouteNotFound))
			return
		}

		if _, ok := errors.AsType[*openapi3filter.SecurityRequirementsError](err); ok {
			logger.Error("security requirements error")

			respondError(w, mapper.ErrorToResp(appErrors.ErrUnauthorized))
			return
		}

		//nolint:nestif // todo: refactor
		if reqErr, ok := errors.AsType[*openapi3filter.RequestError](err); ok {
			logger.Error("request error")

			if reqErr.RequestBody != nil && reqErr.RequestBody.Required && r.ContentLength == 0 {
				resp := mapper.ErrorToResp(appErrors.ErrInvalid)
				resp.Body = &api.ErrorResponse_Body{}
				if err := resp.Body.FromErrorResponseWholeRequestError(
					api.ErrorResponseWholeRequestErrorRequired,
				); err != nil {
					logger.WithError(err).Error("failed to write request error response")
					respondError(w, mapper.ErrorToResp(appErrors.ErrUnknown))
					return
				}
				respondError(w, resp)
				return
			}

			//nolint:nestif // todo: refactor
			if childErr := reqErr.Unwrap(); childErr != nil {
				if schemaErr, ok := errors.AsType[*openapi3.SchemaError](err); ok {
					resp := mapper.ErrorToResp(appErrors.ErrInvalid)
					path := strings.Join(schemaErr.JSONPointer(), ".")

					resp.Body = &api.ErrorResponse_Body{}
					if path != "" {
						fieldError := "invalid"
						if schemaErr.SchemaField == "required" {
							fieldError = "required"
						}
						if err := resp.Body.FromErrorResponseRequestFieldError(
							map[string]string{path: fieldError},
						); err != nil {
							logger.WithError(err).Error("failed to write schema error response")
							respondError(w, mapper.ErrorToResp(appErrors.ErrUnknown))
							return
						}
					} else {
						if err := resp.Body.FromErrorResponseWholeRequestError(
							api.ErrorResponseWholeRequestErrorInvalid,
						); err != nil {
							logger.WithError(err).Error("failed to write schema error response")
							respondError(w, mapper.ErrorToResp(appErrors.ErrUnknown))
							return
						}
					}

					respondError(w, resp)
					return
				}
				if _, ok := errors.AsType[*openapi3filter.ParseError](err); ok {
					p := reqErr.Parameter

					resp := mapper.ErrorToResp(appErrors.ErrInvalid)
					if p.In == "path" {
						resp.UrlParams = &api.ErrorResponse_UrlParams{}
						if err := resp.UrlParams.FromErrorResponseRequestFieldError(api.ErrorResponseRequestFieldError{
							p.Name: "invalid",
						}); err != nil {
							logger.WithError(err).Error("failed to write parser error response")
							respondError(w, mapper.ErrorToResp(appErrors.ErrUnknown))
							return
						}
					}
					// todo: query

					respondError(w, resp)
					return
				}

				respondError(w, mapper.ErrorToResp(childErr))
				return
			}
		}

		respondError(w, mapper.ErrorToResp(appErrors.ErrInternal))
	}
}
