package di

import (
	"github.com/hifat/con-q-api/internal/app/handler"
	"github.com/hifat/con-q-api/internal/app/middleware"
)

type Adapter struct {
	Handler    handler.Handler
	Middleware middleware.Middleware
}

func NewAdapter(h handler.Handler, m middleware.Middleware) Adapter {
	return Adapter{
		Handler:    h,
		Middleware: m,
	}
}
