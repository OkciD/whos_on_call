package mapper

import (
	"errors"
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/shared/errors"

	"github.com/OkciD/whos_on_call/internal/shared/models/api"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/utils"
)

type ErrorResp struct {
	api.ErrorResponse
	StatusCode int
}

func newErrorResp(statusCode int, errCode api.ErrorResponseCode) ErrorResp {
	return ErrorResp{
		StatusCode:    statusCode,
		ErrorResponse: api.ErrorResponse{Code: errCode},
	}
}

var errToRespMap = map[error]ErrorResp{
	appErrors.ErrUnauthorized:       newErrorResp(http.StatusUnauthorized, api.Unauthorized),
	appErrors.ErrEntityNotFound:     newErrorResp(http.StatusNotFound, api.EntityNotFound),
	appErrors.ErrRouteRouteNotFound: newErrorResp(http.StatusNotFound, api.RouteNotFound),

	appErrors.ErrDuplicate: newErrorResp(http.StatusBadRequest, api.Duplicate),
	appErrors.ErrInvalid:   newErrorResp(http.StatusBadRequest, api.Invalid),
}
var defaultErrorResp = newErrorResp(http.StatusInternalServerError, "internal")

func ErrorToResp(err error) ErrorResp {
	for errorCandidate, errorResp := range errToRespMap {
		if errors.Is(err, errorCandidate) {
			return errorResp
		}
	}

	return defaultErrorResp
}

var respToErrMap = utils.ReverseMap(errToRespMap)
var defaultErr = appErrors.ErrUnknown

func RespToError(statusCode int, errResp api.ErrorResponse) error {
	key := newErrorResp(statusCode, errResp.Code)
	if err, ok := respToErrMap[key]; ok {
		return err
	}

	return defaultErr
}
