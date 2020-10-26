package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/alextanhongpin/graphql-server-starter/pkg/database"
	"github.com/alextanhongpin/pkg/grace"
)

// NewDB creates an instance of a production database.
// Throws error when the options are not provided.
func NewDB(sg *grace.ShutdownGroup) *sql.DB {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	sg.Add(func(ctx context.Context) {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	})

	return db
}
