package ydisk

import (
	"context"
	"diskSync/src/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ydiskFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Size int64  `json:"size,omitempty"`
}

type yandexResourceResponse struct {
	Embedded struct {
		Items []ydiskFile `json:"items"`
	} `json:"_embedded"`
}

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) ListFiles(ctx context.Context) ([]ydiskFile, error) {
	apiURL := "https://cloud-api.yandex.net/v1/disk/resources?path=/"

	cfg, err := config.Load("src/config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	// Читаем токен как обычный текст, убирая лишние пробелы и переносы строк
	tokenBytes, err := os.ReadFile(cfg.YandexDisk.TokenPath)
	if err != nil {
		return nil, fmt.Errorf("read token file: %w", err)
	}
	token := strings.TrimSpace(string(tokenBytes))

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "OAuth "+token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("yandex api status: %s (code %d)", resp.Status, resp.StatusCode)
	}

	var rawResponse yandexResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	// Возвращаем чистый слайс элементов
	return rawResponse.Embedded.Items, nil
}
