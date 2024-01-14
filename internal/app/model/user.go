package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	Email     string         `gorm:"type:varchar(100); default:null" json:"email"`
	Username  string         `gorm:"type:varchar(100); unique" json:"username"`
	Password  string         `gorm:"type:varchar(200);" json:"password"`
	Name      string         `gorm:"type:varchar(100);" json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
