package middlewareDomain

import (
	"context"

	"github.com/hifat/con-q-api/internal/pkg/token"
)

type IMiddlewareService interface {
	AuthGuard(ctx context.Context, authToken string) (*token.AuthClaims, error)
}
