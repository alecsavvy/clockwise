package state

import (
	"log/slog"

	"github.com/alecsavvy/clockwise/common"
)

type PeerState struct {
	IsHealthy bool
	NodeType  string
}

// internal state of the application held in memory
type AppState struct {
	logger *slog.Logger
	config *common.Config
}

func New(logger *slog.Logger, config *common.Config) (*AppState, error) {
	return &AppState{
		logger: logger,
		config: config,
	}, nil
}
