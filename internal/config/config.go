package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Proxy string `yaml:"proxy"`
}

func readYaml() {
	configFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		slog.Error("Failed to read config file.", "error", err)
	}
	var config Config
	yaml.Unmarshal(configFile, &config)
	if err != nil {
		slog.Error("Failed to parse config file.", "error", err)
	}
}
