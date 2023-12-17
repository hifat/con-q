package middlewareDomain

import (
	"context"

	"github.com/hifat/con-q-api/internal/pkg/token"
)

type IMiddlewareService interface {
	AuthGuard(ctx context.Context, headerToken string) (*token.AuthClaims, error)
	RefreshTokenGuard(ctx context.Context, headerToken string) (*token.AuthClaims, error)
}
