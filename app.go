package main

import (
	"github.com/cometbft/cometbft/abci/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	e    *echo.Echo
	db   *gorm.DB
	abci *types.Application
}
