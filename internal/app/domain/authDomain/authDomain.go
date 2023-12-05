package authDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
)

type IAuthRepo interface {
	Register(ctx context.Context, req ReqRegister) error
	Count(ctx context.Context, userID uuid.UUID) (int64, error)
	Create(ctx context.Context, req ReqAuth) error
	Delete(ctx context.Context, authID uuid.UUID) error
	RemoveTokenExpires(ctx context.Context, userID uuid.UUID) error
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
	Agent    string `json:"-"`
	ClientIP string `json:"-"`
}

type ReqAuth struct {
	ID        uuid.UUID
	Agent     string
	ClientIP  string
	UserID    uuid.UUID
	ExpiresAt time.Time
}

type Passport struct {
	userDomain.User
}

type ResToken struct {
	ID           uuid.UUID `json:"id" example:"594d7791-d4fc..."`
	AccessToken  string    `json:"accessToken" example:"eyJhbGciO..."`
	RefreshToken string    `json:"refreshToken" example:"eyJhbGciO..."`
}
