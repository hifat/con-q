package authSrv

import (
	"context"

	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
)

// When user login remove token that expired
// Check max device

type authSrv struct {
	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) authDomain.IAuthSrv {
	return &authSrv{authRepo, userRepo}
}

func (s *authSrv) Register(ctx context.Context, req authDomain.ReqRegister) error {
	exists, err := s.userRepo.Exists("username", req.Username)
	if err != nil {
		return ernos.InternalServerError()
	}

	if exists {
		return ernos.HasAlreadyExists("username")
	}

	err = s.authRepo.Register(ctx, req)
	if err != nil {
		return ernos.InternalServerError()
	}

	return nil
}
