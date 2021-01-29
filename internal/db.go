package internal

import (
	"database/sql"
	"log"

	"github.com/alextanhongpin/postgres"
	"github.com/markbates/pkger"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	MigrationSource = &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir("github.com/alextanhongpin/go-graphql-template:/internal/postgres/migrations"),
	}
	MigrationsSource     = postgres.WithMigrationsSource(MigrationSource)
	TestMigrationsSource = postgres.WithTestMigrationsSource(MigrationSource)
)

// NewDB creates an instance of a production database.
// Throws error when the options are not provided.
func NewDB() (*sql.DB, func()) {
	cs := postgres.NewConnString()
	db, err := postgres.New(cs.String(), MigrationsSource)
	if err != nil {
		log.Fatal(err)
	}

	return db, func() {
		_ = db.Close()
	}
}
