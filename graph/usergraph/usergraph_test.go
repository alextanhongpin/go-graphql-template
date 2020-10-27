// +build integration

package usergraph_test

import (
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"

	"github.com/alextanhongpin/graphql-server-starter/app"
	"github.com/alextanhongpin/graphql-server-starter/pkg/database"
)

var pool *redis.Pool

func TestMain(m *testing.M) {
	_, closeDB := database.SetupTestDB()
	cache, closeCache := app.SetupTestCache()
	pool = cache

	code := m.Run()

	closeDB()
	closeCache()
	os.Exit(code)
}
