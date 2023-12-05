package model

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	Agent     string    `gorm:"type:varchar(100)" json:"agent"`
	ClientIP  string    `gorm:"type:text" json:"clientIP"`
	ExpiresAt time.Time `json:"expiresAt"`

	UserID uuid.UUID `gorm:"type:uuid" json:"userID"`
	User   User      `json:"user"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
