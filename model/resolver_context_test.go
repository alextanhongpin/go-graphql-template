package model_test

import (
	"testing"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"
)

type repo struct {
	entity.Querier
}

func TestResolverContextFactory(t *testing.T) {
	resolver := model.NewResolverContext(&repo{})
	if resolver == nil {
		t.Fatal("expected resolver not to be nil, got nil")
	}
}
