package model

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	Id        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	Agent     string    `gorm:"type:varchar(100)" json:"agent"`
	ClientIP  string    `gorm:"type:varchar(30)" json:"clientIP"`
	ExpiresAt time.Time `json:"expiresAt"`

	UserId uuid.UUID `gorm:"type:uuid" json:"userId"`
	User   User      `json:"user"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
