package authMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/middlewareDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
	"github.com/hifat/con-q-api/internal/pkg/helper"
)

type AuthMiddleware struct {
	cfg config.AppConfig

	middlewareService middlewareDomain.IMiddlewareService
}

func New(cfg config.AppConfig, middlewareService middlewareDomain.IMiddlewareService) AuthMiddleware {
	return AuthMiddleware{cfg, middlewareService}
}

func (m *AuthMiddleware) AuthGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			httpResponse.Error(ctx, errorDomain.Error{
				Status:  http.StatusBadRequest,
				Message: authConst.Msg.NO_AUTHORIZATION_HEADER,
				Code:    authConst.Code.NO_AUTHORIZATION_HEADER,
			})
			return
		}

		authClaims, err := m.middlewareService.AuthGuard(ctx, authHeader)
		if err != nil {
			httpResponse.Error(ctx, err)
			return
		}

		ctx.Set("credentials", authClaims)
		ctx.Next()
	}
}

func (m *AuthMiddleware) RefreshTokenGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.Request.Header.Get("X-Refresh-Token")
		if headerToken == "" {
			httpResponse.Error(ctx, errorDomain.Error{
				Status:  http.StatusBadRequest,
				Message: authConst.Msg.NO_X_REFRESH_TOKEN_HEADER,
				Code:    authConst.Code.NO_X_REFRESH_TOKEN_HEADER,
			})
			return
		}

		authClaims, err := m.middlewareService.RefreshTokenGuard(ctx, headerToken)
		if err != nil {
			httpResponse.Error(ctx, err)
			return
		}

		ctx.Set("credentials", authClaims)
		ctx.Set("refreshToken", helper.GetBearerToken(headerToken))
		ctx.Next()
	}
}
