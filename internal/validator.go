package internal

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type Validator interface {
	Validate(ctx context.Context, js interface{}) error
}

type JsonValidator struct {
	v *validator.Validate
}

func NewValidator() *JsonValidator {
	return &JsonValidator{
		v: validator.New(),
	}
}

func (v JsonValidator) Validate(ctx context.Context, js interface{}) error {
	if err := conform.Strings(js); err != nil {
		return err
	}
	return v.v.Struct(js)
}
