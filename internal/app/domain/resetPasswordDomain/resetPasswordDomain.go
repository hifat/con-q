package resetPasswordDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
)

type IResetPasswordRepo interface {
	FirstByCol(ctx context.Context, col string, expected any) (*ResetPassword, error)
	Exists(ctx context.Context, resetId uuid.UUID) (bool, error)
	Create(ctx context.Context, req ReqCreate) error
	CanUsed(ctx context.Context, resetId uuid.UUID) (bool, error)
	MakeUsed(ctx context.Context, resetId uuid.UUID) error
	RevokedByCol(ctx context.Context, col string, expected any) error
}

type IResetPasswordService interface {
	Request(ctx context.Context, req ReqCreate) (*httpDomain.ResSucces[any], error)
	Reset(ctx context.Context, req ReqResetPassword) (*httpDomain.ResSucces[any], error)
}

type ResetPassword struct {
	Id        uuid.UUID  `json:"id"`
	UserId    uuid.UUID  `json:"userId"`
	Code      string     `json:"code"`
	Agent     string     `json:"agent"`
	ClientIP  string     `json:"clientIP"`
	UsedAt    *time.Time `json:"usedAt"`
	RevokedAt *time.Time `json:"revokedAt"`
	ExpiresAt time.Time  `json:"expiresAt"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type ReqCreate struct {
	Email     string     `binding:"required" json:"email"`
	Id        *uuid.UUID `json:"-"`
	Code      string     `json:"-"`
	UserId    uuid.UUID  `json:"-"`
	Agent     string     `json:"-"`
	ClientIP  string     `json:"-"`
	ExpiresAt time.Time  `json:"-"`
}

type ReqResetPassword struct {
	Password string `binding:"required" json:"password"`
	Code     string `binding:"required" json:"code"`
}
