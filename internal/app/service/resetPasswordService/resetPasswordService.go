package resetPasswordService

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/constant/resetPasswordConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/resetPasswordDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/mailer"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

type resetPasswordService struct {
	cfg config.AppConfig

	resetPasswordRepo resetPasswordDomain.IResetPasswordRepo
	userRepo          userDomain.IUserRepo
}

func New(
	cfg config.AppConfig,
	resetPasswordRepo resetPasswordDomain.IResetPasswordRepo,
	userRepo userDomain.IUserRepo,
) resetPasswordDomain.IResetPasswordService {
	return &resetPasswordService{
		cfg,
		resetPasswordRepo,
		userRepo,
	}
}

func (s *resetPasswordService) Request(ctx context.Context, req resetPasswordDomain.ReqCreate) error {
	var user userDomain.User
	err := s.userRepo.FirstByCol(ctx, &user, "email", req.Email)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			return ernos.RecordNotFound("user")
		}

		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	err = s.resetPasswordRepo.RevokedByCol(ctx, "user_id", user.ID)
	if err != nil {
		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	newID := uuid.New()
	code := strings.Split(newID.String(), "-")[0]
	req = resetPasswordDomain.ReqCreate{
		ID:        &newID,
		Email:     user.Email,
		Code:      code,
		UserID:    user.ID,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		ExpiresAt: time.Now().Add(s.cfg.Auth.RESET_PASSWORD_DURATION),
	}

	err = s.resetPasswordRepo.Create(ctx, req)
	if err != nil {
		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	reqSendEmail := mailer.ReqSendEmail{
		From:       "contact@conq.com",
		To:         user.Email,
		TemplateId: "",
		TemplateModel: map[string]string{
			code: code,
		},
	}

	sendEmail := func() {
		_, err := mailer.SendEmail(reqSendEmail)
		if err != nil {
			zlog.Error(err)
		}
	}

	go sendEmail()

	return nil
}

func (s *resetPasswordService) Reset(ctx context.Context, req resetPasswordDomain.ReqResetPassword) error {
	reset, err := s.resetPasswordRepo.FirstByCol(ctx, "code", req.Code)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			return ernos.RecordNotFound("reset password request")
		}

		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	canUsed, err := s.resetPasswordRepo.CanUsed(ctx, reset.ID)
	if err != nil {
		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	if !canUsed {
		return errorDomain.Error{
			Status:  http.StatusBadRequest,
			Message: resetPasswordConst.Msg.CAN_NOT_USED,
			Code:    resetPasswordConst.Code.CAN_NOT_USED,
		}
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	reqResetPassword := userDomain.ReqUpdatePassword{
		Password: hashedPassword,
	}
	err = s.userRepo.UpdatePassword(ctx, reset.UserID, reqResetPassword)
	if err != nil {
		zlog.Error(err)
		return errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		}
	}

	makeUsed := func() {
		err := s.resetPasswordRepo.MakeUsed(ctx, reset.ID)
		zlog.Error(err)
	}
	go makeUsed()

	return nil
}
