package main

import (
	"context"
	"diskSync/src/internal/config"
	"diskSync/src/internal/differ"
	"diskSync/src/internal/gdrive"
	"diskSync/src/internal/lib/logger"
	"diskSync/src/internal/localfs"
	"diskSync/src/internal/storage"
	"diskSync/src/internal/ydisk"
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
	clientGoogle, err := gdrive.New(ctx, config.GoogleDrive.ClientID, config.GoogleDrive.ClientSecret, config.GoogleDrive.TokenPath)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("successful connection", slog.String("connected service", "google drive"))
	clientYandex, err := ydisk.New(ctx, config.YandexDisk.ClientID, config.YandexDisk.ClientSecret, config.YandexDisk.TokenPath)
	if err != nil {
		log.Error("error", slog.String("error", err.Error()))
		return
	}
	log.Info("successful connection", slog.String("connected service", "yandex disk"))
	clientLocal := localfs.New(config.LocalDir)
	storages := map[string]storage.Storage{
		"Google Drive": clientGoogle,
		"Yandex Disk":  clientYandex,
		"Local Files":  clientLocal,
	}
	storageLocal, err := storages["Local Files"].ListFiles(ctx)
	storageGoogle, err := storages["Google Drive"].ListFiles(ctx)
	diff := differ.Diff(storageLocal, storageGoogle)
	log.Info("Files to sync", slog.String("storage", "Google Drive"), slog.String("files to sunc", fmt.Sprint(len(diff))))

}
