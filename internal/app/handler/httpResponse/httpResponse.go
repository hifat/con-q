package httpResponse

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/pkg/validity"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

var cfg *config.AppConfig

func init() {
	cfg = config.LoadAppConfig()
}

func Error(ctx *gin.Context, err any) {
	if e, ok := err.(errorDomain.Error); ok {
		ctx.AbortWithStatusJSON(e.Status, errorDomain.Response{
			Error: e,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorDomain.Response{
		Error: errorDomain.Error{
			Status:  http.StatusInternalServerError,
			Message: commonConst.Msg.INTERNAL_SERVER_ERROR,
			Code:    commonConst.Code.INTERNAL_SERVER_ERROR,
		},
	})
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
	ctx.AbortWithStatusJSON(http.StatusBadRequest, errorDomain.Response{
		Error: errorDomain.Error{
			Status:  http.StatusBadRequest,
			Message: err.(error).Error(),
			Code:    commonConst.Code.BAD_REQUEST,
		},
	})
}

func ValidateFormErr(ctx *gin.Context, err error) {
	validate, err := validity.Validate(err)
	if err != nil {
		zlog.Error(err)
		Error(ctx, err)
		return
	}

	formErr(ctx, validate)
}

func formErr(ctx *gin.Context, err any) {
	resError := handleErr(err)
	if cfg.Env.AppMode == "develop" {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, resError)
		return
	}

	jsonBytes, err := json.MarshalIndent(resError, "", "  ")
	if err != nil {
		Error(ctx, err)
		return
	}

	zlog.Skip(0).Warn(string(jsonBytes))
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errorDomain.Response{
		Error: errorDomain.Error{
			Status:  http.StatusUnprocessableEntity,
			Message: commonConst.Msg.INVALID_FORM_VALIDATION,
			Code:    commonConst.Code.INVALID_FORM_VALIDATION,
		},
	})
}

func Created(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusCreated, obj)
}

func Success(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}
