package authRepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) authDomain.IAuthRepo {
	return &authRepo{db}
}

func (r *authRepo) Exists(ctx context.Context, authId uuid.UUID) (bool, error) {
	var exists bool
	return exists, r.db.WithContext(ctx).
		Model(&model.Auth{}).
		Select("COUNT(*) > 0").
		Where("id = ?", authId).
		Find(&exists).Error
}

func (r *authRepo) Register(ctx context.Context, req authDomain.ReqRegister) error {
	var newUser model.User
	err := copier.Copy(&newUser, &req)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Create(&newUser).Error
}

func (r *authRepo) Count(ctx context.Context, userId uuid.UUID) (int64, error) {
	var count int64
	return count, r.db.WithContext(ctx).
		Model(&model.Auth{}).
		Where("user_id", userId).
		Count(&count).Error
}

func (r *authRepo) Save(ctx context.Context, req authDomain.ReqAuth) error {
	var newUser model.Auth
	err := copier.Copy(&newUser, &req)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Save(&newUser).Error
}

func (r *authRepo) Delete(ctx context.Context, authId uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", authId).
		Delete(&model.Auth{}).Error
}

func (r *authRepo) RemoveTokenExpires(ctx context.Context, userId uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("expires_at <= ?", time.Now().Format(time.RFC3339)).
		Delete(&model.Auth{}).Error
}
