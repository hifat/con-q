package middlewareSrv

import (
	"context"
	"errors"
	"net/http"

	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/middlewareDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

type middlewareSrv struct {
	cfg config.AppConfig

	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(cfg config.AppConfig, authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) middlewareDomain.IMiddlewareService {
	return &middlewareSrv{
		cfg,
		authRepo,
		userRepo,
	}
}

func (s *middlewareSrv) AuthGuard(ctx context.Context, authToken string) (*token.AuthClaims, error) {
	claims, err := token.Claims(s.cfg.Auth, token.ACCESS, authToken)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrInvalidToken):
			return nil, ernos.Other(errorDomain.Error{
				Status:  http.StatusUnauthorized,
				Message: authConst.Msg.INVALID_TOKEN,
				Code:    authConst.Code.INVALID_TOKEN,
			})
		case errors.Is(err, token.ErrTokenExpired):
			return nil, ernos.Other(errorDomain.Error{
				Status:  http.StatusUnauthorized,
				Message: authConst.Msg.TOKEN_EXPIRED,
				Code:    authConst.Code.TOKEN_EXPIRED,
			})
		default:
			zlog.Error(err)
			return nil, ernos.InternalServerError()
		}
	}

	return claims, nil
}
