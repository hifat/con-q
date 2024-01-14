package authService

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/constant/authConst"
	"github.com/hifat/con-q-api/internal/app/constant/commonConst"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/domain/errorDomain"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/helper"
	"github.com/hifat/con-q-api/internal/pkg/token"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	cfg config.AppConfig

	authRepo authDomain.IAuthRepo
	userRepo userDomain.IUserRepo
}

func New(cfg config.AppConfig, authRepo authDomain.IAuthRepo, userRepo userDomain.IUserRepo) authDomain.IAuthService {
	return &authService{cfg, authRepo, userRepo}
}

func (s *authService) Register(ctx context.Context, req authDomain.ReqRegister) (*httpDomain.ResSucces[any], error) {
	exists, err := s.userRepo.Exists("username", req.Username)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	if exists {
		return nil, ernos.HasAlreadyExists("username")
	}

	req.Password, err = helper.HashPassword(req.Password)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = s.authRepo.Register(ctx, req)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	res := &httpDomain.ResSucces[any]{
		Message: "ok",
	}

	return res, nil
}

func (s *authService) Login(ctx context.Context, req authDomain.ReqLogin) (*httpDomain.ResSucces[authDomain.ResToken], error) {
	var user userDomain.User
	err := s.userRepo.FirstByCol(ctx, &user, "username", req.Username)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			return nil, ernos.InvalidCredentials()
		}

		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, ernos.InvalidCredentials()
	}

	err = s.authRepo.RemoveTokenExpires(ctx, user.Id)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	count, err := s.authRepo.Count(ctx, user.Id)
	if count >= int64(s.cfg.Auth.MaxDevice) {
		return nil, ernos.Other(errorDomain.Error{
			Status:  http.StatusForbidden,
			Message: authConst.Msg.MAX_DEVICES_LOGIN,
			Code:    authConst.Code.MAX_DEVICES_LOGIN,
		})
	}

	passport := &authDomain.Passport{
		User: userDomain.User{
			Id:       user.Id,
			Username: user.Username,
			Name:     user.Name,
		},
	}

	passport.AuthId = uuid.New()
	resToken, exp, err := s.generateToken(*passport)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = s.authRepo.Save(ctx, authDomain.ReqAuth{
		Id:        passport.AuthId,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserId:    user.Id,
		ExpiresAt: exp.Refresh,
	})
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	res := &httpDomain.ResSucces[authDomain.ResToken]{
		Item: resToken,
	}

	return res, nil
}

func (s *authService) Logout(ctx context.Context, tokenId uuid.UUID) (*httpDomain.ResSucces[any], error) {
	err := s.authRepo.Delete(ctx, tokenId)
	if err != nil {
		return nil, ernos.InternalServerError()
	}

	res := &httpDomain.ResSucces[any]{
		Message: "ok",
	}

	return res, err
}

func (s *authService) RefreshToken(ctx context.Context, passport authDomain.Passport, req authDomain.ReqRefreshToken) (*httpDomain.ResSucces[authDomain.ResToken], error) {
	var user userDomain.User
	err := s.userRepo.FirstByCol(ctx, &user, "username", passport.User.Username)
	if err != nil {
		if err.Error() == commonConst.Msg.RECORD_NOTFOUND {
			zlog.Error(err)
			return nil, ernos.InvalidCredentials()
		}

		return nil, ernos.InternalServerError()
	}

	claims, err := token.Claims(s.cfg.Auth, token.REFRESH, req.RefreshToken)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	authId := claims.Passport.AuthId
	exists, err := s.authRepo.Exists(ctx, authId)
	if !exists {
		return nil, ernos.RevokedToken()
	}

	err = s.authRepo.Delete(ctx, authId)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	passport.AuthId = uuid.New()
	resToken, exp, err := s.generateToken(passport)
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	err = s.authRepo.Save(ctx, authDomain.ReqAuth{
		Id:        passport.AuthId,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserId:    user.Id,
		ExpiresAt: exp.Refresh,
	})
	if err != nil {
		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	res := &httpDomain.ResSucces[authDomain.ResToken]{
		Item: resToken,
	}

	return res, nil
}
