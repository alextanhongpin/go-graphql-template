package model

import "github.com/alextanhongpin/graphql-server-starter/entity"

type ResolverContext struct {
	Repository entity.Querier
}

func NewResolverContext(r entity.Querier) *ResolverContext {
	return &ResolverContext{
		Repository: r,
	}
}
