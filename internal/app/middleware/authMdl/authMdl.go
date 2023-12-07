package authMdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/middlewareDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
)

type AuthMiddleware struct {
	cfg config.AppConfig

	middlewareSrv middlewareDomain.IMiddlewareService
}

func New(cfg config.AppConfig, middlewareSrv middlewareDomain.IMiddlewareService) AuthMiddleware {
	return AuthMiddleware{cfg, middlewareSrv}
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

		authClaims, err := m.middlewareSrv.AuthGuard(ctx, authHeader)
		if err != nil {
			httpResponse.Error(ctx, err)
			return
		}

		ctx.Set("credentials", authClaims)
		ctx.Next()
	}
}
