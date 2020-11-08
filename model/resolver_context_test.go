package model_test

import (
	"testing"

	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/model"
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
