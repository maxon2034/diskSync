package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	switch env {
	case "local":
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		return logger
	case "dev", "prod":
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		return logger
	default:
		return nil
	}
}
