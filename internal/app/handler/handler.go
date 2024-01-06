package handler

import (
	"github.com/hifat/con-q-api/internal/app/handler/authHandler"
	"github.com/hifat/con-q-api/internal/app/handler/healtzHandler"
	"github.com/hifat/con-q-api/internal/app/handler/resetPasswordHandler"
	"github.com/hifat/con-q-api/internal/app/handler/userHandler"
)

type Handler struct {
	Healtz        healtzHandler.HealtzHandler
	Auth          authHandler.AuthHandler
	ResetPassword resetPasswordHandler.ResetPasswordHandler
	User          userHandler.UserHandler
}

func New(
	HealtzHandler healtzHandler.HealtzHandler,
	AuthHandler authHandler.AuthHandler,
	ResetPasswordHandler resetPasswordHandler.ResetPasswordHandler,
	UserHandler userHandler.UserHandler,
) Handler {
	return Handler{
		Healtz:        HealtzHandler,
		Auth:          AuthHandler,
		ResetPassword: ResetPasswordHandler,
		User:          UserHandler,
	}
}
