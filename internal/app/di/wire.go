//go:build wireinject
// +build wireinject

package di

import (
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/database"
	"github.com/hifat/con-q-api/internal/app/handler"
	"github.com/hifat/con-q-api/internal/app/handler/authHdl"
	"github.com/hifat/con-q-api/internal/app/handler/healtzHdl"
	"github.com/hifat/con-q-api/internal/app/repository/authRepo"
	"github.com/hifat/con-q-api/internal/app/repository/userRepo"
	"github.com/hifat/con-q-api/internal/app/service/authSrv"

	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	database.NewPostgresConnection,
	authRepo.New,
	userRepo.New,
)

var ServiceSet = wire.NewSet(
	authSrv.New,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	healtzHdl.New,
	authHdl.New,
)

func InitializeAPI(cfg config.AppConfig) (Adapter, func()) {
	wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	return Adapter{}, nil
}
