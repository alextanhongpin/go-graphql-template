// +build integration

package resolver_test

import (
	"testing"

	"github.com/gomodule/redigo/redis"

	"github.com/alextanhongpin/postgres"
)

var pool *redis.Pool

func TestMain(m *testing.M) {
	_, stop := postgres.InitTestDB(internal.TestMigrationsSource)
	defer stop()

	var close func()
	pool, close = internal.InitTestCache()
	defer close()

	m.Run()
}
