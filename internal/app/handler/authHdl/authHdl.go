package authHdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
	"github.com/hifat/con-q-api/internal/pkg/validity"
)

type AuthHandler struct{}

func New() AuthHandler {
	return AuthHandler{}
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
