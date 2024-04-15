package app

import (
	"log/slog"
	"sync"

	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/peer"
	"github.com/alecsavvy/clockwise/server"
	"github.com/alecsavvy/clockwise/storage"
)

type App struct {
	/** common */
	config *common.Config
	state  *common.AppState
	logger *slog.Logger

	/** services */
	rpcService     *server.RpcService
	peerService    *peer.PeerService
	storageService *storage.StorageService
}

func New(logger *slog.Logger, server *server.RpcService, db *storage.StorageService, discovery *peer.PeerService) (*App, error) {
	return &App{
		logger:         logger,
		rpcService:     server,
		storageService: db,
		peerService:    discovery,
	}, nil
}

func (app *App) Run() error {
	type taskFunc func() error
	tasks := []taskFunc{
		app.peerService.ConnectPeers,
		app.peerService.PollPeerHealth,
		app.rpcService.Serve,
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t taskFunc) {
			defer wg.Done()
			err := t()
			if err != nil {
				app.logger.Error("task crashed:", err)
			}
		}(task)
	}

	wg.Wait()
	return nil
}
