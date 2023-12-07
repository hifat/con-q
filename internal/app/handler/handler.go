package handler

import (
	"github.com/hifat/con-q-api/internal/app/handler/authHdl"
	"github.com/hifat/con-q-api/internal/app/handler/healtzHdl"
)

type Handler struct {
	Healtz healtzHdl.HealtzHandler
	Auth   authHdl.AuthHandler
}

func New(HealtzHandler healtzHdl.HealtzHandler, AuthHandler authHdl.AuthHandler) Handler {
	return Handler{
		Healtz: HealtzHandler,
		Auth:   AuthHandler,
	}
}
