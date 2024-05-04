package main

import (
	"github.com/alecsavvy/clockwise/db"
	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/node"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *echo.Echo
	db   *db.DB
	abci *types.Application
	n    *node.Node
}

func NewApp() (*App, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.HideBanner = true

	return &App{
		e:  e,
		db: db,
	}, nil
}
