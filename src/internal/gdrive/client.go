package gdrive

import (
	"context"
	"fmt"

	"google.golang.org/api/drive/v3"
)

type Client struct {
	Service *drive.Service
}

func (c *Client) ListFiles(ctx context.Context) ([]*drive.File, error) {
	result, err := c.Service.Files.List().PageSize(100).Fields("files(id, name, mimeType)").Do()
	if err != nil {
		return nil, fmt.Errorf("Error in listing files: %w", err)
	}
	return result.Files, nil
}
