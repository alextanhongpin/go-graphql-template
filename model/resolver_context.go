package model

import (
	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/pkg/loader"
)

type ResolverContext struct {
	Repository entity.Querier
	UserLoader loader.UserLoader
}

func NewResolverContext(r entity.Querier) *ResolverContext {
	return &ResolverContext{
		Repository: r,
		UserLoader: loader.NewUser(r),
	}
}
