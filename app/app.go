package app

import (
	"log/slog"
	"sync"

	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/peer"
	"github.com/alecsavvy/clockwise/server"
)

// entrypoint to most logic connecting the discovery, hashrings, dbs, and stuff together
type App struct {
	logger *slog.Logger
	server *server.Server
	db     *db.DB
	peers  *peer.PeerManager
}

func New(logger *slog.Logger, server *server.Server, db *db.DB, discovery *peer.PeerManager) (*App, error) {
	return &App{
		logger: logger,
		server: server,
		db:     db,
		peers:  discovery,
	}, nil
}

func (app *App) Run() error {
	type taskFunc func() error
	tasks := []taskFunc{
		app.peers.ConnectPeers,
		app.peers.PollPeerHealth,
		app.server.Serve,
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
