package errors

import (
	"errors"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/utils"
)

type ErrorResp struct {
	StatusCode int
	ErrCode    string
}

func newErrorResp(statusCode int, errCode string) ErrorResp {
	return ErrorResp{
		StatusCode: statusCode,
		ErrCode:    errCode,
	}
}

var errToRespMap = map[error]ErrorResp{
	ErrUnauthorized:   newErrorResp(http.StatusUnauthorized, "unauthorized"),
	ErrNotFound:       newErrorResp(http.StatusNotFound, "not found"),
	ErrNotImplemented: newErrorResp(http.StatusNotImplemented, "not impl"),
	ErrBadJSON:        newErrorResp(http.StatusBadRequest, "bad json"),

	ErrDuplicate: newErrorResp(http.StatusBadRequest, "duplicate"),
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
var defaultErr = ErrUnknown

func RespToError(errResp ErrorResp) error {
	if err, ok := respToErrMap[errResp]; ok {
		return err
	}

	return defaultErr
}
