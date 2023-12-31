package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResetPassword struct {
	ID        uuid.UUID  `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	Code      string     `gorm:"type:varchar(20)" json:"code"`
	Agent     string     `gorm:"type:varchar(100)" json:"agent"`
	ClientIP  string     `gorm:"type:varchar(30)" json:"clientIP"`
	UsedAt    *time.Time `json:"usedAt"`
	RevokedAt *time.Time `json:"revokedAt"`
	ExpiresAt time.Time  `json:"expiresAt"`

	UserID uuid.UUID `gorm:"type:uuid" json:"userID"`
	User   User      `json:"user"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
