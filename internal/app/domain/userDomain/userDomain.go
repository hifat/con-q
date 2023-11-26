package userDomain

import "time"

type IUserRepo interface {
	Exists(col string, expected string) (bool, error)
}

type User struct {
	Username  string     `json:"username" example:"conq"`
	Name      string     `json:"name" example:"Corn Dog"`
	CreatedAt *time.Time `json:"createdAt" example:"2023-11-24T13:00:00Z"`
}
