package authHdl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
	"github.com/hifat/con-q-api/internal/pkg/validity"
)

type AuthHandler struct {
	authSrv authDomain.IAuthSrv
}

func New(authSrv authDomain.IAuthSrv) AuthHandler {
	return AuthHandler{authSrv}
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
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	err = h.authSrv.Register(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err.Error())
		e := err.(errorDomain.Error)

		ctx.JSON(e.Status, e)
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
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	var res authDomain.ResToken
	err = h.authSrv.Login(ctx.Request.Context(), &res, req)
	if err != nil {
		log.Println(err.Error())
		e := err.(errorDomain.Error)

		ctx.JSON(e.Status, e)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
