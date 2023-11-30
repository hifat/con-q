package authDomain

import (
	"context"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
)

type IAuthRepo interface {
	Register(ctx context.Context, req ReqRegister) error
}

type IAuthSrv interface {
	Register(ctx context.Context, req ReqRegister) error
	Login(ctx context.Context, req ReqLogin) (*ResToken, error)
	Logout(ctx context.Context, tokenID uuid.UUID) error
}

type ReqRegister struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`
	Password string `binding:"required,min=8,max=70" json:"password" example:"Cq123456_"`
	Name     string `binding:"required,max=100" json:"name" example:"Corn Dog"`
}

type ReqLogin struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`
	Password string `binding:"required,min=8,max=70" json:"password" example:"Cq123456_"`
}

type Passport struct {
	userDomain.User
}

type ResToken struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciO..."`
	RefreshToken string `json:"refreshToken" example:"eyJhbGciO..."`
}
