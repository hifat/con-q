package userDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type IUserRepo interface {
	Exists(col string, expected string) (bool, error)
	FirstByCol(ctx context.Context, user *User, col string, expected any) error
	UpdatePassword(ctx context.Context, userID uuid.UUID, req ReqUpdatePassword) error
}

type User struct {
	ID        uuid.UUID  `json:"userID" example:"60cf8c94-2c98..."`
	Email     string     `json:"email" example:"conq@domain.com"`
	Username  string     `json:"username" example:"conq"`
	Name      string     `json:"name" example:"Corn Dog"`
	Password  string     `json:"-"`
	CreatedAt *time.Time `json:"createdAt,omitempty" example:"2023-11-24T13:00:00Z"`
	UpdatedAt *time.Time `json:"udpatedAt,omitempty" example:"2023-11-24T13:00:00Z"`
}

type ReqUpdatePassword struct {
	Password string `json:"password"`
}
