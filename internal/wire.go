package internal

import "github.com/google/wire"

var Set = wire.NewSet(
	NewRedisConfig,
	NewRedis,
	NewDB,
	NewLoggerConfig,
	NewLogger,
	NewAuthorizerConfig,
	NewAuthorizer,
)
