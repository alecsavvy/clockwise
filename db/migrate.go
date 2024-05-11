package db

import (
	"database/sql"
	"embed"

	"github.com/alecsavvy/clockwise/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed sql/migrations/*
var dbMigrations embed.FS

func RunMigrations(pgConnectionString string) error {
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
