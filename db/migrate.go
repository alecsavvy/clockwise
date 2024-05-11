package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alecsavvy/clockwise/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func RunMigrations(logger *utils.Logger, pgConnectionString string) error {
	tries := 10
	for {
		if tries < 0 {
			return errors.New("ran out of retries for migrations")
		}
		err := runMigrations(logger, pgConnectionString)
		if err != nil {
			tries = tries - 1
			logger.Info("issue running migrations", "error", err, "tries_left", tries)
			time.Sleep(3 * time.Second)
			continue
		}
		return nil
	}
}

func runMigrations(logger *utils.Logger, pgConnectionString string) error {
	db, err := sql.Open("pgx", pgConnectionString)
	if err != nil {
		return utils.AppError("error opening sql db", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return utils.AppError("could not ping db", err)
	}

	migrationsDir := "./sql/migrations"

	goose.SetLogger(logger)

	err = goose.SetDialect("postgres")
	if err != nil {
		return utils.AppError("error setting goose dialect", err)
	}

	err = goose.Up(db, migrationsDir)
	if err != nil {
		return utils.AppError("error on goose up", err)
	}

	return nil
}
