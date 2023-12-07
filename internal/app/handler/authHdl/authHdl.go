package authHdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
	"github.com/hifat/con-q-api/internal/pkg/validity"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

type AuthHandler struct {
	cfg config.AppConfig

	authSrv authDomain.IAuthSrv
}

func New(cfg config.AppConfig, authSrv authDomain.IAuthSrv) AuthHandler {
	return AuthHandler{cfg, authSrv}
}

// @Summary		Register
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ReqRegister
// @Success		409 {object} errorDomain.Response "Duplicate record"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/register [post]
// @Param		Body body authDomain.ReqRegister true "Register request"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req authDomain.ReqRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		validate, err := validity.Validate(err)
		if err != nil {
			zlog.Error(err)
			httpResponse.Error(ctx, err)
			return
		}

		httpResponse.FormErr(ctx, h.cfg.Env.AppMode, validate)
		return
	}

	err = h.authSrv.Register(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// @Summary		Login
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResToken
// @Success		401 {object} errorDomain.Response "Invalid Credentials"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/login [post]
// @Param		Body body authDomain.ReqLogin true "Login request"
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req authDomain.ReqLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		validate, err := validity.Validate(err)
		if err != nil {
			zlog.Error(err)
			httpResponse.Error(ctx, validate)
			return
		}

		httpResponse.FormErr(ctx, h.cfg.Env.AppMode, validate)
		return
	}

	req.Agent = ctx.Request.UserAgent()
	req.ClientIP = ctx.ClientIP()

	res, err := h.authSrv.Login(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, httpDomain.ResSucces{
		Item: res,
	})
}

// @Summary		Refresh Token
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResToken
// @Success		401 {object} errorDomain.Response "Invalid Credentials"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/auth/refresh-token [post]
// @Param		Body body authDomain.ReqLogin true "Login request"
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var req authDomain.ReqLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		validate, err := validity.Validate(err)
		if err != nil {
			zlog.Error(err)
			httpResponse.Error(ctx, validate)
			return
		}

		httpResponse.FormErr(ctx, h.cfg.Env.AppMode, validate)
		return
	}

	req.Agent = ctx.Request.UserAgent()
	req.ClientIP = ctx.ClientIP()

	res, err := h.authSrv.Login(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, httpDomain.ResSucces{
		Item: res,
	})
}
