package userRepo

import (
	"context"

	"github.com/google/uuid"
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
		Where(map[string]interface{}{col: expected}).Find(&exists).Error
}

func (r *userRepo) FirstByCol(ctx context.Context, user *userDomain.User, col string, expected any) (err error) {
	return r.db.Model(&model.User{}).
		Where(map[string]any{col: expected}).
		First(&user).Error
}

func (r *userRepo) UpdatePassword(ctx context.Context, userID uuid.UUID, req userDomain.ReqUpdatePassword) error {
	return r.db.Model(&model.User{
		ID: userID,
	}).Updates(model.User{
		Password: req.Password,
	}).Error
}
