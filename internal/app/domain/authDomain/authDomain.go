package authDomain

import "context"

type IAuthRepo interface {
	Register(ctx context.Context, req ReqRegister) error
}

type IAuthSrv interface {
	Register(ctx context.Context, req ReqRegister) error
}

type ReqRegister struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`         // Your username
	Password string `binding:"required,min=8,max=75" json:"password" example:"Cq1234_"` // Your password
	Name     string `binding:"required,max=100" json:"name" example:"Corn Dog"`         // Your full name
}
