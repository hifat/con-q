package httpResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
)

func Error(ctx *gin.Context, err any) {
	if e, ok := err.(errorDomain.Error); ok {
		ctx.AbortWithStatusJSON(e.Status, errorDomain.Response{
			Error: e,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func handleErr(err any) errorDomain.Response {
	if _, ok := err.(error); ok {
		return errorDomain.Response{
			Error: errorDomain.Error{
				Message: err.(error).Error(),
			},
		}
	}

	return errorDomain.Response{
		Error: errorDomain.Error{
			Attribute: err,
		},
	}
}

func BadRequest(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err.(error).Error())
}

func FormErr(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, handleErr(err))
}

func Created(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusCreated, obj)
}

func Success(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}
