//go:build wireinject
// +build wireinject

package di

import (
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/database"
	"github.com/hifat/con-q-api/internal/app/handler"
	"github.com/hifat/con-q-api/internal/app/handler/authHandler"
	"github.com/hifat/con-q-api/internal/app/handler/healtzHandler"
	"github.com/hifat/con-q-api/internal/app/handler/resetPasswordHandler"
	"github.com/hifat/con-q-api/internal/app/middleware"
	"github.com/hifat/con-q-api/internal/app/middleware/authMiddleware"
	"github.com/hifat/con-q-api/internal/app/repository/authRepo"
	"github.com/hifat/con-q-api/internal/app/repository/resetPasswordRepo"
	"github.com/hifat/con-q-api/internal/app/repository/userRepo"
	"github.com/hifat/con-q-api/internal/app/service/authService"
	"github.com/hifat/con-q-api/internal/app/service/middlewareService"
	"github.com/hifat/con-q-api/internal/app/service/resetPasswordService"

	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	database.NewPostgresConnection,
	authRepo.New,
	userRepo.New,
	resetPasswordRepo.New,
)

var ServiceSet = wire.NewSet(
	authService.New,
	middlewareService.New,
	resetPasswordService.New,
)

var HandlerSet = wire.NewSet(
	handler.New,
	middleware.New,

	authMiddleware.New,
	healtzHandler.New,
	authHandler.New,
	resetPasswordHandler.New,
)

var AdapterSet = wire.NewSet(NewAdapter)

func InitializeAPI(cfg config.AppConfig) (Adapter, func()) {
	wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	return Adapter{}, nil
}
