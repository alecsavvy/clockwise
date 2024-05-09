package utils

import (
	"log/slog"
	"os"

	"github.com/cometbft/cometbft/libs/log"
)

type Logger struct {
	log slog.Logger
}

func NewLogger(opts *slog.HandlerOptions) *Logger {
	logger := *slog.New(slog.NewJSONHandler(os.Stdout, opts))
	return &Logger{
		logger,
	}
}

func (l *Logger) Debug(msg string, keyvals ...interface{}) {
	l.log.Debug(msg, keyvals...)
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	l.log.Info(msg, keyvals...)
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	l.log.Error(msg, keyvals...)
}

func (l *Logger) With(keyvals ...interface{}) log.Logger {
	newLogger := l.log.With(keyvals...)
	return &Logger{log: *newLogger}
}

var _ log.Logger = (*Logger)(nil)
