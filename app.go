package main

import (
	"log/slog"

	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/peer"
	"github.com/alecsavvy/clockwise/server"
	"github.com/alecsavvy/clockwise/state"
	"github.com/alecsavvy/clockwise/storage"
	"github.com/alecsavvy/clockwise/ui"
)

type App struct {
	/** common */
	config *common.Config
	logger *slog.Logger
	state  *state.AppState

	/** services */
	rpcService     *server.RpcService
	peerService    *peer.PeerService
	storageService *storage.StorageService
	uiService      *ui.UIService
}

func NewApp(plogger *slog.Logger, config *common.Config) (*App, error) {
	logger := plogger.With("module", "app")

	// initialize app state
	state, err := state.New(logger, config)
	if err != nil {
		return nil, err
	}

	return &App{
		logger: logger,
		config: config,
		state:  state,
	}, nil
}

func (app *App) Run() error {
	common.Await(app.logger, app.peerService.Run, app.rpcService.Run, app.storageService.Run, app.uiService.Run)
	return nil
}
