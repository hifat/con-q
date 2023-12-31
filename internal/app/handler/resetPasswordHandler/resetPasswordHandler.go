package resetPasswordHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/resetPasswordDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
)

type ResetPasswordHandler struct {
	resetPasswordService resetPasswordDomain.IResetPasswordService
}

func New(resetPasswordService resetPasswordDomain.IResetPasswordService) ResetPasswordHandler {
	return ResetPasswordHandler{resetPasswordService}
}

// @Summary		Request Reset Password
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} resetPasswordDomain.ReqCreate
// @Success		409 {object} errorDomain.Response "Duplicate record"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/reset-password [post]
// @Param		Body body resetPasswordDomain.ReqCreate true "Request request"
func (h *ResetPasswordHandler) Request(ctx *gin.Context) {
	var req resetPasswordDomain.ReqCreate
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.ValidateFormErr(ctx, err)
		return
	}

	req.ClientIP = ctx.ClientIP()
	req.Agent = ctx.Request.UserAgent()

	err = h.resetPasswordService.Request(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// @Summary		Reset Password
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} resetPasswordDomain.ReqResetPassword
// @Success		409 {object} errorDomain.Response "Duplicate record"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/reset-password [patch]
// @Param		Body body resetPasswordDomain.ReqResetPassword true "Request request"
func (h *ResetPasswordHandler) Reset(ctx *gin.Context) {
	var req resetPasswordDomain.ReqResetPassword
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.ValidateFormErr(ctx, err)
		return
	}

	err = h.resetPasswordService.Reset(ctx.Request.Context(), req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
