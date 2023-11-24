package authRepo

import (
	"context"

	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) authDomain.IAuthRepo {
	return &authRepo{db}
}

func (r *authRepo) Register(ctx context.Context, req authDomain.ReqRegister) error {
	return r.db.WithContext(ctx).
		Create(&model.User{
			Username: req.Username,
			Password: req.Password,
			Name:     req.Name,
		}).Error
}
