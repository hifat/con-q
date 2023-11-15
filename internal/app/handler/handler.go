package handler

import (
	"github.com/hifat/con-q/internal/app/handler/healtzHdl"
)

type Handler struct {
	Healtz healtzHdl.HealtzHandler
}

func NewHandler(HealtzHandler healtzHdl.HealtzHandler) Handler {
	return Handler{
		Healtz: HealtzHandler,
	}
}
