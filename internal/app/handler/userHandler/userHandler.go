package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/commonDomain"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/app/handler/httpResponse"
)

type UserHandler struct {
	userService userDomain.IUserService
}

func New(userService userDomain.IUserService) UserHandler {
	return UserHandler{
		userService,
	}
}

// @Summary		Get Users
// @Tags		User
// @Accept		json
// @Produce		json
// @Success		200 {object} resetPasswordDomain.ReqCreate
// @Success		409 {object} errorDomain.Response "Duplicate record"
// @Success		422 {object} errorDomain.Response "Form validation error"
// @Success		500 {object} errorDomain.Response "Internal server error"
// @Router		/reset-password [post]
// @Param		Body body resetPasswordDomain.ReqCreate true "Request request"
func (h *UserHandler) Get(ctx *gin.Context) {
	var reqQuery commonDomain.ReqQuery
	err := ctx.ShouldBindQuery(&reqQuery)
	if err != nil {
		httpResponse.ValidateFormErr(ctx, err)
		return
	}

	users, err := h.userService.Get(ctx.Request.Context(), reqQuery)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	total := 0
	ctx.JSON(http.StatusOK, httpDomain.ResSucces{
		Items: users,
		Meta: &httpDomain.ResMeta{
			Total:   &total,
			Page:    reqQuery.Page,
			PerPage: reqQuery.PerPage,
		},
	})
}
