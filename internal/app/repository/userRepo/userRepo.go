package userRepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/commonDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"github.com/hifat/con-q-api/internal/app/repository"
	"gorm.io/gorm"
)

var (
	fields = []string{"email", "username", "password", "name", "createdAt", "updatedAt"}
)

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) userDomain.IUserRepo {
	return &userRepo{db}
}

func (r *userRepo) Get(ctx context.Context, query commonDomain.ReqQuery) (users []userDomain.User, total int64, err error) {
	tx := r.db.Model(&model.User{})

	reqQuery := repository.NewQueryRequest(tx, fields, query)
	err = reqQuery.Validate()
	if err != nil {
		return users, total, err
	}
	reqQuery.Search(repository.CONTAIN, repository.OR)
	reqQuery.Sort()

	tx.Count(&total)

	reqQuery.Pagination()

	return users, total, tx.Find(&users).Error
}

func (r *userRepo) Exists(col string, expected string) (exists bool, err error) {
	return exists, r.db.Model(&model.User{}).
		Select("COUNT(*) > 0").
		Where(map[string]interface{}{col: expected}).
		Find(&exists).Error
}

func (r *userRepo) FirstByCol(ctx context.Context, user *userDomain.User, col string, expected any) (err error) {
	return r.db.Model(&model.User{}).
		Where(map[string]any{col: expected}).
		First(&user).Error
}

func (r *userRepo) UpdatePassword(ctx context.Context, userId uuid.UUID, req userDomain.ReqUpdatePassword) error {
	return r.db.Model(&model.User{
		Id: userId,
	}).Updates(model.User{
		Password: req.Password,
	}).Error
}
