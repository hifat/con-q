package resetPasswordRepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/resetPasswordDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"gorm.io/gorm"
)

type resetPasswordRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) resetPasswordDomain.IResetPasswordRepo {
	return &resetPasswordRepo{db}
}

func (r *resetPasswordRepo) FirstByCol(ctx context.Context, col string, expected any) (*resetPasswordDomain.ResetPassword, error) {
	var res resetPasswordDomain.ResetPassword
	return &res, r.db.Model(&model.ResetPassword{}).
		Where(map[string]any{
			col: expected,
		}).
		First(&res).Error
}

func (r *resetPasswordRepo) Exists(ctx context.Context, resetID uuid.UUID) (bool, error) {
	var exists bool
	return exists, r.db.Model(&model.ResetPassword{}).
		Select("COUNT(*) > 0").
		Where("id = ?", resetID).
		Find(&exists).Error
}

func (r *resetPasswordRepo) Create(ctx context.Context, req resetPasswordDomain.ReqCreate) error {
	newID := uuid.New()
	if req.ID != nil {
		newID = *req.ID
	}

	return r.db.Create(&model.ResetPassword{
		ID:        newID,
		Code:      req.Code,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		ExpiresAt: req.ExpiresAt,
		UserID:    req.UserID,
	}).Error
}

func (r *resetPasswordRepo) CanUsed(ctx context.Context, resetID uuid.UUID) (bool, error) {
	var canUsed bool
	return canUsed, r.db.Model(&model.ResetPassword{}).
		Select("COUNT(*) > 0").
		Where("id = ?", resetID).
		Where("expires_at > ?", time.Now().Format(time.RFC3339)).
		Find(&canUsed).Error
}

func (r *resetPasswordRepo) MakeUsed(ctx context.Context, resetID uuid.UUID) error {
	return r.db.Model(&model.ResetPassword{
		ID: resetID,
	}).Updates(model.ResetPassword{
		IsUsed: true,
	}).Error
}
