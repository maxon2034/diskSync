package localfs

import (
	"context"
	"diskSync/src/internal/storage"
	"fmt"
	"os"
)

type Client struct {
	rootDir string
}

func New(dir string) *Client {
	return &Client{rootDir: dir}
}

func (c *Client) ListFiles(ctx context.Context) ([]storage.File, error) {
	entry, err := os.ReadDir(c.rootDir)
	if err != nil {
		return nil, fmt.Errorf("Error in reading directory: %v", err)
	}
	storage := make([]storage.File, len(entry))
	for i, v := range entry {
		storage[i].Name = v.Name()
		if v.IsDir() == true {
			storage[i].Type = "dir"
		} else {
			storage[i].Type = "file"
		}
	}
	return storage, nil
}
