package repository

import "errors"

var (
	SomeFieldsNotFoundErr = errors.New("some fields not found")
	FieldNotFoundErr      = errors.New("field not found")
	InvalidSortTypeErr    = errors.New("sort type must be 'ASC' or 'DESC'")
	InvalidSortFormatErr  = errors.New("sort must be format `field:sortType'")
	OperatorNotFoundErr   = errors.New("operator not found")
)

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}
