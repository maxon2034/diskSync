package gdrive

import (
	"context"
	"diskSync/src/internal/storage"
	"fmt"

	"google.golang.org/api/drive/v3"
)

type Client struct {
	Service *drive.Service
}

func (c *Client) ListFiles(ctx context.Context) ([]storage.File, error) {
	storage := make([]storage.File, 100)
	result, err := c.Service.Files.List().PageSize(100).Fields("files(id, name, mimeType)").Do()
	if err != nil {
		return nil, fmt.Errorf("Error in listing files: %w", err)
	}
	if len(result.Files) == 0 {
		return nil, fmt.Errorf("no files on google drive")
	}
	for i, v := range result.Files {
		storage[i].Name = v.Name
		storage[i].Type = v.Kind
	}
	return storage, nil
}
