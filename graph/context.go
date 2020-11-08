package graph

import (
	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/go-playground/validator"
)

type Context struct {
	Repository entity.Querier
	Validator  *validator.Validate
}

func NewContext(r entity.Querier) *Context {
	return &Context{
		Repository: r,
		Validator:  validator.New(),
	}
}
