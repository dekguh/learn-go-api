package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

type Config struct {
	Application AppConfig `yaml:"application"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = "configs/development/app.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
