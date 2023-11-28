package authSrv

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"golang.org/x/crypto/bcrypt"
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

	req.Password, err = helper.HashPassword(req.Password)
	if err != nil {
		return ernos.InternalServerError()
	}

	err = s.authRepo.Register(ctx, req)
	if err != nil {
		return ernos.InternalServerError()
	}

	return nil
}

func (s *authSrv) Login(ctx context.Context, res *authDomain.ResToken, req authDomain.ReqLogin) error {
	var user userDomain.User
	err := s.userRepo.FirstByCol(ctx, &user, "username", req.Username)
	if err != nil {
		return ernos.InvalidCredentials()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return ernos.InvalidCredentials()
	}

	cfg := config.LoadAppConfig()
	claims := &authDomain.Passport{
		User: userDomain.User{
			Username: user.Username,
			Name:     user.Name,
		},
	}

	newToken := token.New(*cfg, user.Password, *claims)

	accessToken, err := newToken.Signed(token.ACCESS)
	if err != nil {
		log.Println(err.Error())
		return ernos.InternalServerError()
	}

	refreshToken, err := newToken.Signed(token.REFRESH)
	if err != nil {
		log.Println(err.Error())
		return ernos.InternalServerError()
	}

	*res = authDomain.ResToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return nil
}

func (s *authSrv) Logout(ctx context.Context, tokenID uuid.UUID) error {
	return nil
}
