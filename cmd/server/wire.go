//+build wireinject

package main

import (
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain"
	"github.com/alextanhongpin/go-graphql-template/external"
	"github.com/alextanhongpin/go-graphql-template/external/security"
	"github.com/alextanhongpin/go-graphql-template/internal"
	"github.com/google/wire"
)

func New() (http.Handler, func()) {
	panic(wire.Build(
		// Build internal.
		internal.Set,

		// Build app.
		domain.Set,

		// Build external.
		security.Set,
		external.NewGraph,
	))
}
