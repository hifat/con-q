//go:build wireinject
// +build wireinject

package di

import (
	"github.com/hifat/con-q/internal/app/config"
	"github.com/hifat/con-q/internal/app/database"
	"github.com/hifat/con-q/internal/app/handler"
	"github.com/hifat/con-q/internal/app/handler/healtzHdl"

	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	database.NewPostgresConnection,
)

// var ServiceSet = wire.NewSet()

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	healtzHdl.New,
)

func InitializeAPI(cfg *config.AppConfig) (Adapter, func()) {
	// wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	wire.Build(AdapterSet, RepoSet, HandlerSet)
	return Adapter{}, nil
}
