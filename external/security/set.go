package security

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAuthorizerConfig,
	NewAuthorizer,
	wire.Bind(new(Authorizer), new(*JwtAuthorizer)),
)
