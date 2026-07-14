package main

import (
	"diskSync/src/internal/config"
	"diskSync/src/internal/lib/logger"
	"log/slog"
)

func main() {
	env, err := config.Load("src/config/config.yaml")
	logger := logger.New(env.Env)
	if err != nil {
		logger.Error("error", slog.String("error", err.Error()))
		return
	}
	logger.Info("config loaded", slog.String("path", "config.yaml"))
}
