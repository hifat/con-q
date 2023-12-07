package middleware

import "github.com/hifat/con-q-api/internal/app/middleware/authMdl"

type Middleware struct {
	Auth authMdl.AuthMiddleware
}

func New(authMdl authMdl.AuthMiddleware) Middleware {
	return Middleware{
		Auth: authMdl,
	}
}
