package internal

import (
	"database/sql"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewDB,
	wire.Bind(new(entity.DBTX), new(*sql.DB)),
	entity.New,
	wire.Bind(new(entity.Querier), new(*entity.Queries)),
	NewLoggerConfig,
	NewLogger,
	NewRedisConfig,
	NewRedis,
	NewValidator,
	wire.Bind(new(Validator), new(*JsonValidator)),
)
