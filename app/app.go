package app

import (
	"context"

	"github.com/alecsavvy/clockwise/config"
	"github.com/alecsavvy/clockwise/service"
)

var _ service.Service = (*App)(nil)

type App struct {
	service.BaseService

	config *config.Config
}

func New() *App {
	return &App{}
}

func (app *App) Name() string {
	return "openaudiod"
}

func (app *App) Init(context.Context) error {
	// config := config.ReadConfig()

	// create db connections, run migrations?

	// pass down config and deps to core, storage, etc services

	// call .Init() on all downstream services

	return nil
}

func (app *App) Start(ctx context.Context) error {
	// app.Run(core.Run)
	// app.run(storage.Run)
	return nil
}
