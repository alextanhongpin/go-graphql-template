package model

import (
	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/pkg/loader"
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
