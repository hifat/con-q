// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hifat/con-q/internal/app/config"
	"github.com/hifat/con-q/internal/app/handler"
	"github.com/hifat/con-q/internal/app/handler/healtzHdl"
)

// Injectors from wire.go:

func InitializeAPI(cfg *config.AppConfig) (Adapter, func()) {
	healtzHandler := healtzHdl.New()
	handlerHandler := handler.NewHandler(healtzHandler)
	adapter := NewAdapter(handlerHandler)
	return adapter, func() {
	}
}

// wire.go:

var HandlerSet = wire.NewSet(handler.NewHandler, healtzHdl.New)
