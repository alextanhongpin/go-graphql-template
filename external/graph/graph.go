package graph

import (
	"github.com/alextanhongpin/go-graphql-template/external/graph/mutation"
	"github.com/alextanhongpin/go-graphql-template/external/graph/query"
)

type Root struct {
	*mutation.Mutation
	*query.Query
}

func NewRoot() *Root {
	return &Root{
		Mutation: mutation.New(),
		Query:    query.New(),
	}
}
