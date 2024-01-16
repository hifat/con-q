package authDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
)

type IAuthRepo interface {
	Exists(ctx context.Context, authID uuid.UUID) (bool, error)
	Register(ctx context.Context, req ReqRegister) error
	Count(ctx context.Context, userID uuid.UUID) (int64, error)
	Save(ctx context.Context, req ReqAuth) error
	Delete(ctx context.Context, authID uuid.UUID) error
	RemoveTokenExpires(ctx context.Context, userID uuid.UUID) error
}

type IAuthService interface {
	Register(ctx context.Context, req ReqRegister) (*httpDomain.ResSucces[any], error)
	Login(ctx context.Context, req ReqLogin) (*httpDomain.ResSucces[ResToken], error)
	Logout(ctx context.Context, tokenID uuid.UUID) (*httpDomain.ResSucces[any], error)
	RefreshToken(ctx context.Context, passport Passport, req ReqRefreshToken) (*httpDomain.ResSucces[ResToken], error)
}

type ReqRegister struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`
	Password string `binding:"required,min=8,max=70" json:"password" example:"Cq123456_"`
	Name     string `binding:"required,max=100" json:"name" example:"Corn Dog"`
}

type ReqLogin struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`
	Password string `binding:"required,min=8,max=70" json:"password" example:"Cq123456_"`
	Agent    string `json:"-"`
	ClientIP string `json:"-"`
}

type ReqAuth struct {
	ID        uuid.UUID
	Agent     string
	ClientIP  string
	ExpiresAt time.Time
	UserID    uuid.UUID
}

type Passport struct {
	AuthID uuid.UUID
	User   userDomain.User
}

type ResToken struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciO..."`
	RefreshToken string `json:"refreshToken" example:"eyJhbGciO..."`
}

type ReqRefreshToken struct {
	RefreshToken string `binding:"required" json:"refreshToken" example:"eyJhbGciO..."`
	Agent        string `json:"-"`
	ClientIP     string `json:"-"`
}
