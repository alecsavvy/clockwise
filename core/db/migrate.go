package db

import (
	"database/sql"
	"errors"
	"time"

	"embed"

	"github.com/alecsavvy/clockwise/utils"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql/migrations/*
var migrationsFS embed.FS

func RunMigrations(logger *utils.Logger, pgConnectionString string) error {
	tries := 10
	db, err := sql.Open("postgres", pgConnectionString)
	if err != nil {
		return utils.AppError("error opening sql db", err)
	}
	defer db.Close()
	for {
		if tries < 0 {
			return errors.New("ran out of retries for migrations")
		}
		err = db.Ping()
		if err != nil {
			tries = tries - 1
			time.Sleep(2 * time.Second)
			continue
		}
		err := runMigrations(logger, db)
		if err != nil {
			logger.Info("issue running migrations", "error", err, "tries_left", tries)
			return utils.AppError("can't run migrations", err)
		}
		return nil
	}
}

func runMigrations(logger *utils.Logger, db *sql.DB) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "sql/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return utils.AppError("error running migrations", err)
	}

	logger.Infof("Applied %d successful migrations!", n)

	return nil
}
