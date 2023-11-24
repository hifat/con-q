package ernos

import (
	"net/http"
	"strings"

	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
)

func HasAlreadyExists(value ...string) error {
	msg := commonConst.Code.DUPLICATE_RECORD
	if len(value) > 0 {
		msg = strings.Join(value, "") + " has already exists"
	}

	return errorDomain.Error{
		Status:  http.StatusConflict,
		Message: msg,
		Code:    commonConst.Code.DUPLICATE_RECORD,
	}
}

func NotFound(value ...string) error {
	msg := commonConst.Code.RECORD_NOTFOUND
	if len(value) > 0 {
		msg = strings.Join(value, "") + " not found"
	}

	return errorDomain.Error{
		Status:  http.StatusNotFound,
		Message: msg,
		Code:    commonConst.Code.RECORD_NOTFOUND,
	}
}

func Forbidden(value string) error {
	msg := http.StatusText(http.StatusForbidden)
	if value != "" {
		msg = value
	}

	return errorDomain.Error{
		Message: msg,
		Code:    commonConst.Code.RECORD_NOTFOUND,
	}
}

func Unauthorized(value ...string) error {
	msg := commonConst.Code.UNAUTHORIZED
	if len(value) > 0 {
		msg = strings.Join(value, "")
	}

	return errorDomain.Error{
		Status:  http.StatusUnauthorized,
		Message: msg,
		Code:    commonConst.Code.UNAUTHORIZED,
	}
}

func InternalServerError(value ...string) error {
	msg := commonConst.Code.INTERNAL_SERVER_ERROR
	if len(value) > 0 {
		msg = strings.Join(value, "")
	}

	return errorDomain.Error{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
	}
}

func Other(e errorDomain.Error) error {
	return e
}

func Detect(err error) error {
	set := map[any]error{
		commonConst.Msg.RECORD_NOTFOUND: NotFound(),
	}

	if _, ok := set[err.Error()]; !ok {
		return InternalServerError()
	}

	return set[err.Error()]
}
