package repository

import (
	"errors"
	"strings"

	"github.com/hifat/con-q-api/internal/app/domain/commonDomain"
	"gorm.io/gorm"
)

var (
	SomeFieldsNotFoundErr = errors.New("some fields not found")
	FieldNotFoundErr      = errors.New("field not found")
	InvalidSortTypeErr    = errors.New("sort type must be 'ASC' or 'DESC'")
	InvalidSortFormatErr  = errors.New("sort must be format `field:sortType'")
)

type queryRequest struct {
	tx     *gorm.DB
	fields []string
	query  commonDomain.ReqQuery
}

func NewQueryRequest(tx *gorm.DB, fields []string, query commonDomain.ReqQuery) *queryRequest {
	return &queryRequest{
		tx,
		fields,
		query,
	}
}

func (q *queryRequest) validateSortType() error {
	if q.query.Sort != nil {
		types := map[string]struct{}{
			"ASC":  {},
			"DESC": {},
		}

		sorts := strings.Split(*q.query.Sort, ":")
		if len(sorts) != 2 {
			return InvalidSortFormatErr
		}

		if _, exists := types[sorts[1]]; !exists {
			return InvalidSortTypeErr
		}
	}

	return nil
}

func (q *queryRequest) validateFields() error {
	if q.query.Fields != nil {
		fieldSet := make(map[string]struct{}, len(q.fields))
		for _, field := range q.fields {
			fieldSet[field] = struct{}{}
		}

		queryFields := strings.Split(*q.query.Fields, ",")
		for _, queryField := range queryFields {
			if _, exists := fieldSet[queryField]; !exists {
				return SomeFieldsNotFoundErr
			}
		}
	}

	return nil
}

func (q *queryRequest) Validate() error {
	if err := q.validateFields(); err != nil {
		return err
	}

	if err := q.validateSortType(); err != nil {
		return err
	}

	return nil
}

func (q *queryRequest) Pagination() (*gorm.DB, error) {
	perPage := 20
	page := 0
	if q.query.Page != nil && q.query.PerPage != nil {
		page = (*q.query.Page - 1) * *q.query.PerPage
		perPage = *q.query.PerPage
	}

	q.tx.Limit(perPage).Offset(page)
	return q.tx, nil
}

func (q *queryRequest) Sort() (*gorm.DB, error) {
	if q.query.Sort != nil {
		sorts := strings.Split(*q.query.Sort, ":")
		field := sorts[0]
		sortBy := sorts[1]

		q.tx.Order(field + " " + sortBy)
	}

	return q.tx, nil
}
