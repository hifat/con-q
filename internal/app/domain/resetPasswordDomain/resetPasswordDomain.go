package resetPasswordDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type IResetPasswordRepo interface {
	FirstByCol(ctx context.Context, col string, expected any) (*ResetPassword, error)
	Exists(ctx context.Context, resetID uuid.UUID) (bool, error)
	Create(ctx context.Context, req ReqCreate) error
	CanUsed(ctx context.Context, resetID uuid.UUID) (bool, error)
	MakeUsed(ctx context.Context, resetID uuid.UUID) error
}

type IResetPasswordService interface {
	Request(ctx context.Context, req ReqCreate) error
	Reset(ctx context.Context, req ReqResetPassword) error
}

type ResetPassword struct {
	ID        uuid.UUID  `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid" json:"userID"`
	Code      string     `gorm:"type:varchar(20)" json:"code"`
	Agent     string     `gorm:"type:varchar(100)" json:"agent"`
	ClientIP  string     `gorm:"type:varchar(30)" json:"clientIP"`
	IsUsed    bool       `gorm:"default:false" json:"isUsed"`
	ExpiresAt time.Time  `json:"expiresAt"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type ReqCreate struct {
	Email     string     `binding:"required" json:"email"`
	ID        *uuid.UUID `json:"-"`
	Code      string     `json:"-"`
	UserID    uuid.UUID  `json:"-"`
	Agent     string     `json:"-"`
	ClientIP  string     `json:"-"`
	ExpiresAt time.Time  `json:"-"`
}

type ReqResetPassword struct {
	Password string `binding:"required" json:"password"`
	Code     string `binding:"required" json:"code"`
}
