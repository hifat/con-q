package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"ID"`
	Username  string         `gorm:"type:varchar(100);unique" json:"username"`
	Password  string         `gorm:"type:varchar(100);" json:"password"`
	Name      string         `gorm:"type:varchar(100);" json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
