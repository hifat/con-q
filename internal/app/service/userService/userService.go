package userService

import (
	"context"

	"github.com/hifat/con-q-api/internal/app/domain/commonDomain"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
	"github.com/hifat/con-q-api/internal/app/domain/userDomain"
	"github.com/hifat/con-q-api/internal/app/repository"
	"github.com/hifat/con-q-api/internal/pkg/ernos"
	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

type userService struct {
	userRepo userDomain.IUserRepo
}

func New(userRepo userDomain.IUserRepo) userDomain.IUserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) Get(ctx context.Context, query commonDomain.ReqQuery) (*httpDomain.ResSucces[userDomain.User], error) {
	searchBy := "name, username"
	query.SearchBy = &searchBy
	users, total, err := s.userRepo.Get(ctx, query)
	if err != nil {
		if e, ok := err.(repository.Error); ok {
			return nil, ernos.BadRequestError(e.Message)
		}

		zlog.Error(err)
		return nil, ernos.InternalServerError()
	}

	res := &httpDomain.ResSucces[userDomain.User]{
		Items: users,
		Meta: &httpDomain.ResMeta{
			Pagination: &httpDomain.Pagination{
				Total:   &total,
				Page:    query.Page,
				PerPage: query.PerPage,
			},
		},
	}

	return res, nil
}
