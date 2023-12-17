package middleware

import "github.com/hifat/con-q-api/internal/app/middleware/authMiddleware"

type Middleware struct {
	Auth authMiddleware.AuthMiddleware
}

func New(authMiddleware authMiddleware.AuthMiddleware) Middleware {
	return Middleware{
		Auth: authMiddleware,
	}
}
