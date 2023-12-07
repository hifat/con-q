package authSrv

import (
	"context"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
	"golang.org/x/crypto/bcrypt"
)

type authSrv struct {
	cfg config.AppConfig

	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(cfg config.AppConfig, authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) authDomain.IAuthSrv {
	return &authSrv{cfg, authRepo, userRepo}
}

func (s *authSrv) Register(ctx context.Context, req authDomain.ReqRegister) error {
	exists, err := s.userRepo.Exists("username", req.Username)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	if exists {
		return ernos.HasAlreadyExists("username")
	}

	req.Password, err = helper.HashPassword(req.Password)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	err = s.authRepo.Register(ctx, req)
	if err != nil {
		zlog.Error(err)
		return ernos.InternalServerError()
	}

	return nil
}

func (s *authSrv) Login(ctx context.Context, req authDomain.ReqLogin) (res *authDomain.ResToken, err error) {
	var user userDomain.User
	err = s.userRepo.FirstByCol(ctx, &user, "username", req.Username)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			zlog.Error(err)
			return nil, ernos.InvalidCredentials()
		}

		return nil, ernos.InternalServerError()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, ernos.InvalidCredentials()
	}

	claims := &authDomain.Passport{
		User: userDomain.User{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
		},
	}

	err = s.authRepo.RemoveTokenExpires(ctx, user.ID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	tokenID := uuid.New()
	res, exp, err := s.generateToken(tokenID, *claims)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = s.authRepo.Create(ctx, authDomain.ReqAuth{
		ID:        tokenID,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserID:    user.ID,
		ExpiresAt: exp.Refresh,
	})
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}

func (s *authSrv) Logout(ctx context.Context, tokenID uuid.UUID) error {
	return s.authRepo.Delete(ctx, tokenID)
}

func (s *authSrv) RefreshToken(ctx context.Context, passport authDomain.Passport, req authDomain.ReqRefreshToken) (*authDomain.ResToken, error) {
	var user userDomain.User
	err := s.userRepo.FirstByCol(ctx, &user, "username", passport.Username)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			zlog.Error(err)
			return nil, ernos.InternalServerError()
		}

		return nil, ernos.InvalidCredentials()
	}

	claims, err := token.Claims(s.cfg.Auth, token.REFRESH, req.RefreshToken)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	claimsID, err := uuid.Parse(claims.ID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = s.authRepo.Delete(ctx, claimsID)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	tokenID := uuid.New()
	res, exp, err := s.generateToken(tokenID, passport)

	err = s.authRepo.Create(ctx, authDomain.ReqAuth{
		ID:        claimsID,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserID:    user.ID,
		ExpiresAt: exp.Refresh,
	})
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	return res, nil
}
