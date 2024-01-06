package repository

import (
	"fmt"
	"strings"

	"github.com/hifat/con-q-api/internal/app/domain/commonDomain"
	"github.com/stoewer/go-strcase"
	"gorm.io/gorm"
)

type queryRequest struct {
	tx     *gorm.DB
	fields []string
	query  commonDomain.ReqQuery
}

// Basic query that commonly used. Support only main model
func NewQueryRequest(tx *gorm.DB, fields []string, query commonDomain.ReqQuery) *queryRequest {
	return &queryRequest{
		tx,
		fields,
		query,
	}
}

func (q *queryRequest) validateSort() error {
	if q.query.Sort != nil {
		types := map[string]struct{}{
			"ASC":  {},
			"DESC": {},
		}

		sorts := strings.Split(*q.query.Sort, ":")
		if len(sorts) != 2 {
			return InvalidSortFormatErr
		}

		field := sorts[0]
		sortType := sorts[1]

		err := q.checkFieldExists(field)
		if err != nil {
			return err
		}

		if _, exists := types[sortType]; !exists {
			return InvalidSortTypeErr
		}
	}

	return nil
}

func (q *queryRequest) checkFieldExists(reqFields string) error {
	fieldSet := make(map[string]struct{}, len(q.fields))
	for _, field := range q.fields {
		fieldSet[field] = struct{}{}
	}

	queryFields := strings.Split(reqFields, ",")
	for _, queryField := range queryFields {
		if _, exists := fieldSet[strings.Trim(queryField, " ")]; !exists {
			return SomeFieldsNotFoundErr
		}
	}

	return nil
}

func (q *queryRequest) validateFields() error {
	if q.query.Fields != nil {
		return q.checkFieldExists(*q.query.Fields)
	}

	return nil
}

func (q *queryRequest) validateSearch() error {
	if q.query.SearchBy != nil && q.query.Search != nil {
		return q.checkFieldExists(*q.query.SearchBy)
	}

	return nil
}

func (q *queryRequest) Validate() error {
	if err := q.validateFields(); err != nil {
		return Error{
			Message: err.Error(),
		}
	}

	if err := q.validateSort(); err != nil {
		return Error{
			Message: err.Error(),
		}
	}

	if err := q.validateSearch(); err != nil {
		return Error{
			Message: err.Error(),
		}
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
		field := strcase.SnakeCase(sorts[0])
		sortBy := sorts[1]

		q.tx.Order(field + " " + sortBy)
	}

	return q.tx, nil
}

type filterType int8
type operatorType int8

const (
	EQ filterType = iota
	CONTAIN
)

const (
	OR operatorType = iota
	AND
)

func (q *queryRequest) getFilter(s filterType) (string, error) {
	filterSet := map[filterType]string{
		EQ:      "?",
		CONTAIN: "%?%",
	}

	if _, exists := filterSet[s]; !exists {
		return "", OperatorNotFoundErr
	}

	return filterSet[s], nil
}

func (q *queryRequest) Search(s filterType, op operatorType) *gorm.DB {
	if q.query.SearchBy == nil || q.query.Search == nil {
		return q.tx
	}

	andFunc := func(condition string, param string) *gorm.DB {
		return q.tx.Where(condition, param)
	}

	orFunc := func(condition string, param string) *gorm.DB {
		return q.tx.Or(condition, param)
	}

	filter, _ := q.getFilter(s)
	reqFields := strings.Split(*q.query.SearchBy, ",")
	for index, reqField := range reqFields {
		condition := fmt.Sprintf("%s ILIKE ?", reqField)
		param := strings.ReplaceAll(filter, "?", *q.query.Search)
		if op == AND || index == 0 {
			andFunc(condition, param)
			continue
		}

		orFunc(condition, param)
	}

	return q.tx
}
