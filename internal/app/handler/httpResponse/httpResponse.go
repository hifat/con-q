package httpResponse

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
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

func FormErr(ctx *gin.Context, appMode string, err any) {
	resError := handleErr(err)
	if appMode == "develop" {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, resError)
		return
	}

	jsonBytes, err := json.MarshalIndent(resError, "", "  ")
	if err != nil {
		Error(ctx, err)
	}

	zlog.Skip(0).Warn(string(jsonBytes))
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errorDomain.Response{
		Error: errorDomain.Error{
			Status:  http.StatusUnprocessableEntity,
			Message: "validate failed",
			Code:    http.StatusText(http.StatusUnprocessableEntity),
		},
	})
	return
}

func Created(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusCreated, obj)
}

func Success(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}
