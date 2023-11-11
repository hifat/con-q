package model

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type auth struct {
	ID       uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	Token    string    `gorm:"type:text;unique" json:"token"`
	Agent    string    `gorm:"type:varchar(100)" json:"agent"`
	ClientIP net.IP    `gorm:"type:text" json:"clientIP"`

	UserID uuid.UUID `gorm:"type:uuid" json:"userID"`
	User   User      `json:"user"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
