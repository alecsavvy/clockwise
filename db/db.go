package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alecsavvy/clockwise/db/gen"
)

//go:embed sql/schema.sql
var ddl string

type DB struct {
	db      *sql.DB
	Queries *gen.Queries
}

func New() (*DB, error) {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := gen.New(db)
	return &DB{
		db:      db,
		Queries: queries,
	}, nil
}
