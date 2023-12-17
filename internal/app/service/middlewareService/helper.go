package middlewareService

import (
	"context"
	"errors"
	"net/http"

	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

func (s *middlewareService) verifyToken(ctx context.Context, tokenType token.TokenType, headerToken string) (*token.AuthClaims, error) {
	accessToken := helper.GetBearerToken(headerToken)
	claims, err := token.Claims(s.cfg.Auth, tokenType, accessToken)
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
