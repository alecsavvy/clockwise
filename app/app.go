package app

import (
	"log/slog"
	"sync"

	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/discovery"
)

// entrypoint to most logic connecting the discovery, hashrings, dbs, and stuff together
type App struct {
	logger    *slog.Logger
	db        *db.DB
	discovery *discovery.Discovery
}

func New(logger *slog.Logger, db *db.DB, discovery *discovery.Discovery) (*App, error) {
	return &App{
		logger:    logger,
		db:        db,
		discovery: discovery,
	}, nil
}

func (app *App) Run() error {
	bootstrapNodes := []string{"http://localhost:6001", "http://localhost:6002"}

	type taskFunc func() error
	tasks := []taskFunc{
		func() error { return app.discovery.DiscoverNodes(bootstrapNodes) },
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
