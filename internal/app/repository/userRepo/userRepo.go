package userRepo

import (
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) userDomain.IUserRepo {
	return &userRepo{db}
}

func (r *userRepo) Exists(col string, expected string) (exists bool, err error) {
	return exists, r.db.Model(&model.User{}).
		Select("COUNT(*) > 0").
		Where(col+" = ?", expected).Find(&exists).Error
}
