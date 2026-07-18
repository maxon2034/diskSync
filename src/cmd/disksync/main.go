package main

import (
	"context"
	"diskSync/src/internal/config"
	"diskSync/src/internal/gdrive"
	"diskSync/src/internal/lib/logger"
	"diskSync/src/internal/ydisk"
	"log/slog"
)

func main() {
	config, err := config.Load("src/config/config.yaml")
	ctx := context.Background()
	var log *slog.Logger
	if err != nil {
		log = logger.New("local")
	} else {
		log = logger.New(config.Env)
	}
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("config loaded", slog.String("path", "config.yaml"))
	_, err = gdrive.New(ctx, config.GoogleDrive.ClientID, config.GoogleDrive.ClientSecret, config.GoogleDrive.TokenPath)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("successful connection", slog.String("connected service", "google drive"))
	_, err = ydisk.New(ctx, config.YandexDisk.ClientID, config.YandexDisk.ClientSecret, config.YandexDisk.TokenPath)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("successful connection", slog.String("connected service", "yandex disk"))
}
