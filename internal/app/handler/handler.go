package handler

import (
	"github.com/hifat/con-q/internal/app/handler/healtzHdl"

	"github.com/google/wire"
)

var NewHandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	HealtzHandler healtzHdl.HealtzHandler
}

func NewHandler(HealtzHandler healtzHdl.HealtzHandler) Handler {
	return Handler{
		HealtzHandler,
	}
}
