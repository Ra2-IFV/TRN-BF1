package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Proxy string `yaml:"proxy"`
}

func (c Config) readConfig() []byte {
	configData, err := os.ReadFile("config/config.yaml")
	if err != nil {
		slog.Error("Failed to read config file.", "error", err)
		return nil
	}
	return configData
}

func (c Config) ParseConfig() Config {
	err := yaml.Unmarshal(c.readConfig(), c)
	if err != nil {
		slog.Error("Failed to unmarshal config", "error", err)
	}
	return c
}
