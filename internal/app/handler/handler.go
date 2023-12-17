package handler

import (
	"github.com/hifat/con-q-api/internal/app/handler/authHandler"
	"github.com/hifat/con-q-api/internal/app/handler/healtzHandler"
)

type Handler struct {
	Healtz healtzHandler.HealtzHandler
	Auth   authHandler.AuthHandler
}

func New(HealtzHandler healtzHandler.HealtzHandler, AuthHandler authHandler.AuthHandler) Handler {
	return Handler{
		Healtz: HealtzHandler,
		Auth:   AuthHandler,
	}
}
