package common

import (
	"log/slog"
	"os"
)

func NewLogger() (*slog.Logger, error) {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	return logger, nil
}
