package service

import (
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

func ErrDetectNotFound(err error) error {
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			zlog.Error(err)
			return ernos.InvalidCredentials()
		}

		return ernos.InternalServerError()
	}

	return nil
}
