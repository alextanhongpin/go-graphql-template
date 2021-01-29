package internal

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate(ctx context.Context, js interface{}) error {
	if err := conform.Strings(js); err != nil {
		return err
	}

	return v.Struct(js)
}
