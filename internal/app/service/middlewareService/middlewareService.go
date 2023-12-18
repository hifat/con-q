package middlewareService

import (
	"context"

	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/middlewareDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/token"
)

type middlewareService struct {
	cfg config.AppConfig

	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(cfg config.AppConfig, authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) middlewareDomain.IMiddlewareService {
	return &middlewareService{
		cfg,
		authRepo,
		userRepo,
	}
}

func (s *middlewareService) AuthGuard(ctx context.Context, headerToken string) (*token.AuthClaims, error) {
	return s.verifyToken(ctx, token.ACCESS, headerToken)
}

func (s *middlewareService) RefreshTokenGuard(ctx context.Context, headerToken string) (*token.AuthClaims, error) {
	return s.verifyToken(ctx, token.REFRESH, headerToken)
}
