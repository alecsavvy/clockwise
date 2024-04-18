package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type App struct {
	/** config */
	host                   string
	maxBlockHistory        uint64
	maxTransactionPerBlock uint64

	/** state */
	currentBlock uint64
	blocks       []Block
	memPool      []Transaction
	peers        []Peer

	/** utils */
	logger *slog.Logger

	/** services */
	e *echo.Echo
}

func NewApp(host string, initialPeers []string) *App {
	var peers []Peer
	for _, endpoint := range initialPeers {
		peers = append(peers, NewPeer(endpoint))
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("node", host)

	e := echo.New()
	e.HideBanner = true

	return &App{
		host:                   host,
		peers:                  peers,
		maxBlockHistory:        100,
		maxTransactionPerBlock: 10,
		currentBlock:           0,
		blocks:                 make([]Block, 0),
		memPool:                make([]Transaction, 0),
		logger:                 logger,
		e:                      e,
	}
}

func (app *App) Run() error {
	app.logger.Info("started")

	// Start server in a goroutine to allow for graceful shutdown handling
	go func() {
		if err := app.e.Start(app.host); err != nil && err != http.ErrServerClosed {
			app.logger.Error("Error starting server:", err)
		}
	}()

	// Set up channel to receive OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Context with timeout for shutting down the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.logger.Info("Shutting down server...")

	// Attempt to gracefully shut down the server
	if err := app.e.Shutdown(ctx); err != nil {
		app.logger.Error("Error during server shutdown:", err)
		return err
	}

	app.logger.Info("Server gracefully stopped")
	return nil
}
