package utils

import (
	"log/slog"
	"os"

	"github.com/cometbft/cometbft/libs/log"
)

type Logger struct {
	slog.Logger
}

func NewLogger(opts *slog.HandlerOptions) *Logger {
	logger := *slog.New(slog.NewJSONHandler(os.Stdout, opts))
	return &Logger{
		logger,
	}
}

func (l *Logger) With(keyvals ...interface{}) log.Logger {
	return l.With(keyvals...)
}

var _ log.Logger = (*Logger)(nil)
