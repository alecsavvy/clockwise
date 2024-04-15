package storage

import (
	"context"
	"database/sql"
	_ "embed"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/storage/gen"
)

//go:embed sql/schema.sql
var ddl string

type StorageService struct {
	logger  *slog.Logger
	config  *common.Config
	db      *sql.DB
	Queries *gen.Queries
}

func New(plogger *slog.Logger, config *common.Config) (*StorageService, error) {
	logger := plogger.With("module", "storage")
	ctx := context.Background()
	db, err := sql.Open("sqlite3", config.DatabaseFilePath)
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := gen.New(db)
	return &StorageService{
		logger:  logger,
		config:  config,
		db:      db,
		Queries: queries,
	}, nil
}

func (ss *StorageService) Run() error {
	return nil
}
