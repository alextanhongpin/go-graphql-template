package model

import (
	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/pkg/loader"
	"github.com/go-playground/validator"
)

type ResolverContext struct {
	Repository entity.Querier
	UserLoader loader.UserLoader
	Validator  *validator.Validate
}

func NewResolverContext(r entity.Querier) *ResolverContext {
	return &ResolverContext{
		Repository: r,
		UserLoader: loader.NewUser(r),
		Validator:  validator.New(),
	}
}
