package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env string `yaml:"env"`
}

func Load(path string) (*Config, error) {
	config := Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error in reading file: %w", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Error in unmarshalling configuration: %w", err)
	}
	return &config, nil
}
