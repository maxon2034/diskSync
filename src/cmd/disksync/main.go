package main

import (
	"context"
	"diskSync/src/internal/config"
	"diskSync/src/internal/gdrive"
	"diskSync/src/internal/lib/logger"
	"fmt"
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
	client, err := gdrive.New(ctx, config.GoogleDrive.ClientID, config.GoogleDrive.ClientSecret, config.GoogleDrive.TokenPath)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("successful connection", slog.String("token", "created"))
	list, err := client.ListFiles(ctx)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
	}
	fmt.Print(len(list))
	for _, v := range list {
		log.Info("- File", slog.String("name", v.Name))
	}
}
