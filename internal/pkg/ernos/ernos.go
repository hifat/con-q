package ernos

import (
	"net/http"
	"strings"

	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
)

func HasAlreadyExists(fields ...string) error {
	msg := commonConst.Code.DUPLICATE_RECORD
	if len(fields) > 0 {
		msg = strings.Join(fields, "") + " has already exists"
	}

	return errorDomain.Error{
		Status:  http.StatusConflict,
		Message: msg,
		Code:    commonConst.Code.DUPLICATE_RECORD,
	}
}

func RecordNotFound(messages ...string) error {
	msg := commonConst.Msg.RECORD_NOTFOUND
	if len(messages) > 0 {
		msg = strings.Join(messages, "") + " not found"
	}

	return errorDomain.Error{
		Status:  http.StatusNotFound,
		Message: msg,
		Code:    commonConst.Code.RECORD_NOTFOUND,
	}
}

func Forbidden() error {
	return errorDomain.Error{
		Status:  http.StatusForbidden,
		Message: strings.ToLower(http.StatusText(http.StatusForbidden)),
		Code:    commonConst.Code.RECORD_NOTFOUND,
	}
}

func Unauthorized() error {
	return errorDomain.Error{
		Status:  http.StatusUnauthorized,
		Message: authConst.Msg.UNAUTHORIZED,
		Code:    authConst.Code.UNAUTHORIZED,
	}
}

func InvalidCredentials(messages ...string) error {
	msg := authConst.Msg.INVALID_CREDENTIALS
	if len(messages) > 0 {
		msg = strings.Join(messages, "")
	}

	return errorDomain.Error{
		Status:  http.StatusUnauthorized,
		Message: msg,
		Code:    authConst.Code.INVALID_CREDENTIALS,
	}
}

func InternalServerError(messages ...string) error {
	msg := commonConst.Code.INTERNAL_SERVER_ERROR
	if len(messages) > 0 {
		msg = strings.Join(messages, "")
	}

	return errorDomain.Error{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
	}
}

func Detect(err error) error {
	set := map[any]error{
		commonConst.Msg.RECORD_NOTFOUND: RecordNotFound(),
	}

	if _, ok := set[err.Error()]; !ok {
		return InternalServerError()
	}

	return set[err.Error()]
}

func Other(e errorDomain.Error) error {
	return e
}
