package main

import (
	"diskSync/src/internal/config"
	"diskSync/src/internal/lib/logger"
	"log/slog"
)

func main() {
	env, err := config.Load("src/config/config.yaml")
	var log *slog.Logger
	if err != nil {
		log = logger.New("local")
	} else {
		log = logger.New(env.Env)
	}

	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("config loaded", slog.String("path", "config.yaml"))
}
