package model

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	CodeUnauthenticated = "UNAUTHENTICATED"
	CodeBadUserInput    = "BAD_USER_INPUT"
	CodeForbidden       = "FORBIDDEN"
	CodeNotFound        = "NOT_FOUND"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type NotFoundError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	ID      ID     `json:"-"`
}

func NewNotFoundError(entity string, id ID) *NotFoundError {
	return &NotFoundError{
		Code:    CodeNotFound,
		Message: fmt.Sprintf("%s %q does not exist or may have been deleted", entity, id),
	}
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e NotFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":      e.Code,
		"message":   e.Message,
		"exception": map[string]interface{}{},
	}
}

func MaybeNotFound(err error, entity string, id ID) error {
	if errors.Is(err, sql.ErrNoRows) {
		return NewNotFoundError(entity, id)
	}
	return err
}
