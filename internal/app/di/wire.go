//go:build wireinject
// +build wireinject

package di

import (
	"github.com/hifat/con-q/internal/app/handler"
	"github.com/hifat/con-q/internal/app/handler/healtzHdl"

	"github.com/google/wire"
)

// var RepoSet = wire.NewSet()

// var ServiceSet = wire.NewSet()

var HandlerSet = wire.NewSet(
	handler.NewHandlerSet,
	healtzHdl.New,
)

func InitializeAPI() (Adapter, func()) {
	// wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	wire.Build(AdapterSet, HandlerSet)
	return Adapter{}, nil
}
